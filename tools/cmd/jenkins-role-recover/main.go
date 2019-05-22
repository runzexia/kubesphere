package main

import (
	"flag"
	"github.com/gocraft/dbr"
	"github.com/golang/glog"
	"kubesphere.io/kubesphere/pkg/db"
	"kubesphere.io/kubesphere/pkg/gojenkins"
	"kubesphere.io/kubesphere/pkg/models/devops"
	"kubesphere.io/kubesphere/pkg/models/tenant"
	"sync"
	"time"
)

const (
	JenkinsAllUserRoleName = "kubesphere-user"
)

var (
	dsn                  string
	dbClient             *db.Database
	jenkinsAdminAddress  string
	jenkinsAdminUsername string
	jenkinsAdminPassword string
	jenkinsClient        *gojenkins.Jenkins
	defaultEventReceiver = db.EventReceiver{}
)

func init() {
	flag.StringVar(&dsn, "database-connection", "root:password@tcp(192.168.0.4:31029)/devops", "data source name")
	flag.StringVar(&jenkinsAdminAddress, "jks-address", "http://192.168.0.2:30180/", "data source name")
	flag.StringVar(&jenkinsAdminUsername, "jks-username", "admin", "username of jenkins")
	flag.StringVar(&jenkinsAdminPassword, "jks-password", "passw0rd", "password of jenkins")

}

func GetJenkins() *gojenkins.Jenkins {
	jenkins := gojenkins.CreateJenkins(nil, jenkinsAdminAddress, 20, jenkinsAdminUsername, jenkinsAdminPassword)
	return jenkins
}

func main() {
	flag.Parse()
	glog.Info("start recover")
	glog.Info(dsn)
	glog.Info(jenkinsAdminAddress)
	glog.Info(jenkinsAdminUsername)
	glog.Info(jenkinsAdminPassword)
	conn, err := dbr.Open("mysql", dsn+"?parseTime=1&multiStatements=1&charset=utf8mb4&collation=utf8mb4_unicode_ci", &defaultEventReceiver)
	if err != nil {
		glog.Fatal(err)
		panic(err)
	}
	conn.SetMaxIdleConns(100)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(10 * time.Second)
	dbClient = &db.Database{
		Session: conn.NewSession(nil),
	}
	err = dbClient.Ping()
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	glog.Info("init db successful")
	jenkins := GetJenkins()
	glog.Info("start jenkins init")
	jenkins, err = jenkins.Init()
	if err != nil {
		glog.Errorf("failed to connect jenkins, %+v", err)
		panic(err)
	}
	globalRole, err := jenkins.GetGlobalRole(JenkinsAllUserRoleName)
	if err != nil {
		glog.Errorf("failed to get jenkins role, %+v", err)
		panic(err)
	}
	if globalRole == nil {
		_, err := jenkins.AddGlobalRole(JenkinsAllUserRoleName, gojenkins.GlobalPermissionIds{
			GlobalRead: true,
		}, true)
		if err != nil {
			glog.Errorf("failed to create jenkins global role, %+v", err)
			panic(err)
		}
	}
	_, err = jenkins.AddProjectRole(JenkinsAllUserRoleName, "\\n\\s*\\r", gojenkins.ProjectPermissionIds{
		SCMTag: true,
	}, true)
	if err != nil {
		glog.Errorf("failed to create jenkins project role, %+v", err)
		panic(err)
	}
	glog.Info("jenkins init successful")
	jenkinsClient = jenkins
	recoverDevOpsProject()
	recoverDevOpsProjectMember()
	glog.Info("recover all success")

}

func recoverDevOpsProject() {
	projects := make([]*devops.V1DevOpsProject, 0)
	_, err := dbClient.
		Select(devops.V1DevOpsProjectColumns...).
		From(devops.DevOpsProjectTableName).
		Where(db.Eq(devops.StatusColumn, devops.StatusActive)).
		Load(&projects)
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	for _, project := range projects {
		var addRoleCh = make(chan *tenant.DevOpsProjectRoleResponse, 8)
		var addRoleWg sync.WaitGroup
		for role, permission := range devops.JenkinsProjectPermissionMap {
			addRoleWg.Add(1)
			go func(role string, permission gojenkins.ProjectPermissionIds) {
				_, err := jenkinsClient.AddProjectRole(devops.GetProjectRoleName(project.ProjectId, role),
					devops.GetProjectRolePattern(project.ProjectId), permission, true)
				addRoleCh <- &tenant.DevOpsProjectRoleResponse{nil, err}
				addRoleWg.Done()
			}(role, permission)
		}
		for role, permission := range devops.JenkinsPipelinePermissionMap {
			addRoleWg.Add(1)
			go func(role string, permission gojenkins.ProjectPermissionIds) {
				_, err := jenkinsClient.AddProjectRole(devops.GetPipelineRoleName(project.ProjectId, role),
					devops.GetPipelineRolePattern(project.ProjectId), permission, true)
				addRoleCh <- &tenant.DevOpsProjectRoleResponse{nil, err}
				addRoleWg.Done()
			}(role, permission)
		}
		addRoleWg.Wait()
		close(addRoleCh)
		for addRoleResponse := range addRoleCh {
			if addRoleResponse.Err != nil {
				glog.Errorf("%+v", addRoleResponse.Err)
				panic(err)
			}
		}
	}

}

func recoverDevOpsProjectMember() {
	memberShips := make([]*devops.DevOpsProjectMembership, 0)
	_, err := dbClient.
		Select(devops.DevOpsProjectMembershipColumns...).
		From(devops.DevOpsProjectMembershipTableName).
		Where(db.Eq(devops.StatusColumn, devops.StatusActive)).
		Load(&memberShips)
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	globalRole, err := jenkinsClient.GetGlobalRole(JenkinsAllUserRoleName)
	if err != nil {
		glog.Errorf("%+v", err)
		panic(err)
	}
	for _, membership := range memberShips {

		err = globalRole.AssignRole(membership.Username)
		if err != nil {
			glog.Errorf("%+v", err)
			panic(err)
		}

		projectRole, err := jenkinsClient.GetProjectRole(devops.GetProjectRoleName(membership.ProjectId, membership.Role))
		if err != nil {
			glog.Errorf("%+v", err)
			panic(err)
		}
		err = projectRole.AssignRole(membership.Username)
		if err != nil {
			glog.Errorf("%+v", err)
			panic(err)
		}

		pipelineRole, err := jenkinsClient.GetProjectRole(devops.GetPipelineRoleName(membership.ProjectId, membership.Role))
		if err != nil {
			glog.Errorf("%+v", err)
			panic(err)
		}
		err = pipelineRole.AssignRole(membership.Username)
		if err != nil {
			glog.Errorf("%+v", err)
			panic(err)
		}
	}
}
