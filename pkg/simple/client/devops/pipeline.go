package devops

import (
	"net/http"
)

type PageableResponse struct {
	Items      []interface{} `json:"items" description:"paging data"`
	TotalCount int           `json:"total_count" description:"total count"`
}

type PipelineList struct {
	Items []Pipeline `json:"items"`
	Total int        `json:"total_count"`
}

// GetPipeline & SearchPipelines
type Pipeline struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability." `
	Links struct {
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		Scm struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"scm,omitempty"`
		Branches struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"branches,omitempty"`
		Actions struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"actions,omitempty"`
		Runs struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"runs,omitempty"`
		Trends struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"trends,omitempty"`
		Queue struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"queue,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource."`
	Actions         []interface{} `json:"actions,omitempty" description:"the list of all actions."`
	Disabled        interface{}   `json:"disabled,omitempty" description:"disable or not, if disabled, can not do any action."`
	DisplayName     string        `json:"displayName,omitempty" description:"display name"`
	FullDisplayName string        `json:"fullDisplayName,omitempty" description:"full display name"`
	FullName        string        `json:"fullName,omitempty" description:"full name"`
	Name            string        `json:"name,omitempty" description:"name"`
	Organization    string        `json:"organization,omitempty" description:"the name of organization"`
	Parameters      interface{}   `json:"parameters,omitempty" description:"parameters of pipeline, a pipeline can define list of parameters pipeline job expects."`
	Permissions     struct {
		Create    bool `json:"create,omitempty" description:"create action"`
		Configure bool `json:"configure,omitempty" description:"configure action"`
		Read      bool `json:"read,omitempty" description:"read action"`
		Start     bool `json:"start,omitempty" description:"start action"`
		Stop      bool `json:"stop,omitempty" description:"stop action"`
	} `json:"permissions,omitempty" description:"permissions"`
	EstimatedDurationInMillis      int           `json:"estimatedDurationInMillis,omitempty" description:"estimated duration time, unit is millis"`
	NumberOfFolders                int           `json:"numberOfFolders,omitempty" description:"number of folders"`
	NumberOfPipelines              int           `json:"numberOfPipelines,omitempty" description:"number of pipelines"`
	PipelineFolderNames            []interface{} `json:"pipelineFolderNames,omitempty" description:"pipeline folder names"`
	WeatherScore                   int           `json:"weatherScore,omitempty" description:"the score to description the result of pipeline activity"`
	BranchNames                    []string      `json:"branchNames,omitempty" description:"branch names"`
	NumberOfFailingBranches        int           `json:"numberOfFailingBranches,omitempty" description:"number of failing branches"`
	NumberOfFailingPullRequests    int           `json:"numberOfFailingPullRequests,omitempty" description:"number of failing pull requests"`
	NumberOfSuccessfulBranches     int           `json:"numberOfSuccessfulBranches,omitempty" description:"number of successful pull requests"`
	NumberOfSuccessfulPullRequests int           `json:"numberOfSuccessfulPullRequests,omitempty" description:"number of successful pull requests"`
	ScmSource                      struct {
		Class  string      `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		APIURL interface{} `json:"apiUrl,omitempty" description:"api url"`
		ID     string      `json:"id,omitempty" description:"The id of the source configuration management (SCM)."`
	} `json:"scmSource,omitempty"`
	TotalNumberOfBranches     int `json:"totalNumberOfBranches,omitempty" description:"total number of branches"`
	TotalNumberOfPullRequests int `json:"totalNumberOfPullRequests,omitempty" description:"total number of pull requests"`
}

// GetPipeBranchRun & SearchPipelineRuns
type BranchPipelineRun struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		PrevRun struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"prevRun,omitempty"`
		Parent struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"parent,omitempty"`
		Tests struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"tests,omitempty"`
		Nodes struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"nodes,omitempty"`
		Log struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"log,omitempty"`
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		BlueTestSummary struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"blueTestSummary,omitempty"`
		Actions struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"actions,omitempty"`
		Steps struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"steps,omitempty"`
		Artifacts struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"artifacts,omitempty"`
		NextRun struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"nextRun,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Actions          []interface{} `json:"actions,omitempty" description:"the list of all actions"`
	ArtifactsZipFile interface{}   `json:"artifactsZipFile,omitempty" description:"the artifacts zip file"`
	CauseOfBlockage  interface{}   `json:"causeOfBlockage,omitempty" description:"the cause of blockage"`
	Causes           []struct {
		Class            string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		ShortDescription string `json:"shortDescription,omitempty" description:"short description"`
		UserID           string `json:"userId,omitempty" description:"user id"`
		UserName         string `json:"userName,omitempty" description:"user name"`
	} `json:"causes,omitempty"`
	ChangeSet                 []interface{} `json:"changeSet,omitempty" description:"changeset information"`
	Description               interface{}   `json:"description,omitempty" description:"description of resource"`
	DurationInMillis          int           `json:"durationInMillis,omitempty" description:"duration time in millis"`
	EnQueueTime               string        `json:"enQueueTime,omitempty" description:"the time of enter the queue"`
	EndTime                   string        `json:"endTime,omitempty" description:"the time of end"`
	EstimatedDurationInMillis int           `json:"estimatedDurationInMillis,omitempty" description:"estimated duration time, unit is millis"`
	ID                        string        `json:"id,omitempty" description:"id"`
	Name                      interface{}   `json:"name,omitempty" description:"name"`
	Organization              string        `json:"organization,omitempty" description:"the name of organization"`
	Pipeline                  string        `json:"pipeline,omitempty" description:"pipeline name"`
	Replayable                bool          `json:"replayable,omitempty" description:"replayable or not"`
	Result                    string        `json:"result,omitempty" description:"the result of pipeline run. e.g. SUCCESS"`
	RunSummary                string        `json:"runSummary,omitempty" description:"pipeline run summary"`
	StartTime                 string        `json:"startTime,omitempty" description:"the time of start"`
	State                     string        `json:"state,omitempty" description:"run state. e.g. RUNNING"`
	Type                      string        `json:"type,omitempty" description:"source type, such as \"WorkflowRun\""`
	Branch                    struct {
		IsPrimary bool          `json:"isPrimary,omitempty" description:"primary or not"`
		Issues    []interface{} `json:"issues,omitempty" description:"issues"`
		URL       string        `json:"url,omitempty" description:"url"`
	} `json:"branch,omitempty"`
	CommitID    string      `json:"commitId,omitempty" description:"commit id"`
	CommitURL   interface{} `json:"commitUrl,omitempty" description:"commit url "`
	PullRequest interface{} `json:"pullRequest,omitempty" description:"pull request"`
}

// GetBranchPipeRunNodes
type BranchPipelineRunNodes struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		Actions struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"actions,omitempty"`
		Steps struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"steps,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Actions            []interface{} `json:"actions,omitempty" description:"the list of all actions"`
	DisplayDescription interface{}   `json:"displayDescription,omitempty" description:"display description"`
	DisplayName        string        `json:"displayName,omitempty" description:"display name"`
	DurationInMillis   int           `json:"durationInMillis,omitempty" description:"duration time in millis"`
	ID                 string        `json:"id,omitempty" description:"id"`
	Input              *Input        `json:"input,omitempty" description:"the action should user input"`
	Result             string        `json:"result,omitempty" description:"the result of pipeline run. e.g. SUCCESS. e.g. SUCCESS"`
	StartTime          string        `json:"startTime,omitempty" description:"the time of start"`
	State              string        `json:"state,omitempty" description:"run state. e.g. RUNNING"`
	Type               string        `json:"type,omitempty" description:"source type, e.g. \"WorkflowRun\""`
	CauseOfBlockage    interface{}   `json:"causeOfBlockage,omitempty" description:"the cause of blockage"`
	Edges              []struct {
		Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		ID    string `json:"id,omitempty" description:"id"`
		Type  string `json:"type,omitempty" description:"source type"`
	} `json:"edges,omitempty"`
	FirstParent interface{} `json:"firstParent,omitempty" description:"first parent resource"`
	Restartable bool        `json:"restartable,omitempty" description:"restartable or not"`
	Steps       []struct {
		Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		Links struct {
			Self struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"self,omitempty"`
			Actions struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"actions,omitempty"`
		} `json:"_links,omitempty"`
		Actions []struct {
			Class string `json:"_class,omitempty"`
			Links struct {
				Self struct {
					Class string `json:"_class,omitempty"`
					Href  string `json:"href,omitempty"`
				} `json:"self,omitempty"`
			} `json:"_links,omitempty"`
			URLName string `json:"urlName,omitempty"`
		} `json:"actions,omitempty" description:"references the reachable path to this resource"`
		DisplayDescription interface{} `json:"displayDescription,omitempty" description:"display description"`
		DisplayName        string      `json:"displayName,omitempty" description:"display name"`
		DurationInMillis   int         `json:"durationInMillis,omitempty" description:"duration time in millis"`
		ID                 string      `json:"id,omitempty" description:"id"`
		Input              *Input      `json:"input,omitempty" description:"the action should user input"`
		Result             string      `json:"result,omitempty" description:"result"`
		StartTime          string      `json:"startTime,omitempty" description:"the time of start"`
		State              string      `json:"state,omitempty" description:"run state. e.g. RUNNING"`
		Type               string      `json:"type,omitempty" description:"source type"`
	} `json:"steps,omitempty"`
}

// Validate
type Validates struct {
	CredentialID string `json:"credentialId,omitempty" description:"the id of credential"`
}

// GetSCMOrg
type SCMOrg struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		Repositories struct {
			Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
			Href  string `json:"href,omitempty" description:"url in api"`
		} `json:"repositories,omitempty"`
		Self struct {
			Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
			Href  string `json:"href,omitempty" description:"self url in api"`
		} `json:"self,omitempty" description:"scm org self info"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Avatar                      string `json:"avatar,omitempty" description:"the url of organization avatar"`
	JenkinsOrganizationPipeline bool   `json:"jenkinsOrganizationPipeline,omitempty" description:"weather or not already have jenkins pipeline."`
	Name                        string `json:"name,omitempty" description:"organization name"`
}

type SCMServer struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		Self struct {
			Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
			Href  string `json:"href,omitempty" description:"self url in api"`
		} `json:"self,omitempty" description:"scm server self info"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	ID     string `json:"id,omitempty" description:"server id of scm server"`
	Name   string `json:"name,omitempty" description:"name of scm server"`
	ApiURL string `json:"apiUrl,omitempty"  description:"url of scm server"`
}

// GetOrgRepo
type OrgRepo struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Repositories struct {
		Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		Links struct {
			Self struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"self,omitempty"`
		} `json:"_links,omitempty" description:"references the reachable path to this resource"`
		Items []struct {
			Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
			Links struct {
				Self struct {
					Class string `json:"_class,omitempty"`
					Href  string `json:"href,omitempty"`
				} `json:"self,omitempty"`
			} `json:"_links,omitempty" description:"references the reachable path to this resource"`
			DefaultBranch string `json:"defaultBranch,omitempty" description:"default branch"`
			Description   string `json:"description,omitempty" description:"description"`
			Name          string `json:"name,omitempty" description:"name"`
			Permissions   struct {
				Admin bool `json:"admin,omitempty" description:"admin"`
				Push  bool `json:"push,omitempty" description:"push action"`
				Pull  bool `json:"pull,omitempty" description:"pull action"`
			} `json:"permissions,omitempty"`
			Private  bool   `json:"private,omitempty" description:"private or not"`
			FullName string `json:"fullName,omitempty" description:"full name"`
		} `json:"items,omitempty"`
		LastPage interface{} `json:"lastPage,omitempty" description:"last page"`
		NextPage interface{} `json:"nextPage,omitempty" description:"next page"`
		PageSize int         `json:"pageSize,omitempty" description:"page size"`
	} `json:"repositories,omitempty"`
}

// StopPipeline
type StopPipe struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		Parent struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"parent,omitempty"`
		Tests struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"tests,omitempty"`
		Nodes struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"nodes,omitempty"`
		Log struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"log,omitempty"`
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		BlueTestSummary struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"blueTestSummary,omitempty"`
		Actions struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"actions,omitempty"`
		Steps struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"steps,omitempty"`
		Artifacts struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"artifacts,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Actions          []interface{} `json:"actions,omitempty" description:"the list of all actions."`
	ArtifactsZipFile interface{}   `json:"artifactsZipFile,omitempty" description:"the artifacts zip file"`
	CauseOfBlockage  interface{}   `json:"causeOfBlockage,omitempty" description:"the cause of blockage"`
	Causes           []struct {
		Class            string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		ShortDescription string `json:"shortDescription,omitempty" description:"short description"`
	} `json:"causes,omitempty"`
	ChangeSet                 []interface{} `json:"changeSet,omitempty" description:"changeset information"`
	Description               interface{}   `json:"description,omitempty" description:"description"`
	DurationInMillis          int           `json:"durationInMillis,omitempty" description:"duration time in millis"`
	EnQueueTime               string        `json:"enQueueTime,omitempty" description:"the time of enter the queue"`
	EndTime                   string        `json:"endTime,omitempty" description:"the time of end"`
	EstimatedDurationInMillis int           `json:"estimatedDurationInMillis,omitempty" description:"estimated duration time in millis"`
	ID                        string        `json:"id,omitempty" description:"id"`
	Name                      interface{}   `json:"name,omitempty" description:"name"`
	Organization              string        `json:"organization,omitempty" description:"the name of organization"`
	Pipeline                  string        `json:"pipeline,omitempty" description:"pipeline"`
	Replayable                bool          `json:"replayable,omitempty" description:"replayable or not"`
	Result                    string        `json:"result,omitempty" description:"the result of pipeline run. e.g. SUCCESS"`
	RunSummary                string        `json:"runSummary,omitempty" description:"pipeline run summary"`
	StartTime                 string        `json:"startTime,omitempty" description:"the time of start"`
	State                     string        `json:"state,omitempty" description:"run state. e.g. RUNNING"`
	Type                      string        `json:"type,omitempty" description:"type"`
	Branch                    struct {
		IsPrimary bool          `json:"isPrimary,omitempty" description:"primary or not"`
		Issues    []interface{} `json:"issues,omitempty" description:"issues"`
		URL       string        `json:"url,omitempty" description:"url"`
	} `json:"branch,omitempty"`
	CommitID    string      `json:"commitId,omitempty" description:"commit id"`
	CommitURL   interface{} `json:"commitUrl,omitempty" description:"commit url"`
	PullRequest interface{} `json:"pullRequest,omitempty" description:"pull request"`
}

// ReplayPipeline
type ReplayPipe struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		Parent struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"parent,omitempty"`
		Tests struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"tests,omitempty"`
		Log struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"log,omitempty"`
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		BlueTestSummary struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"blueTestSummary,omitempty"`
		Actions struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"actions,omitempty"`
		Artifacts struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"artifacts,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Actions          []interface{} `json:"actions,omitempty" description:"the list of all actions."`
	ArtifactsZipFile interface{}   `json:"artifactsZipFile,omitempty" description:"the artifacts zip file"`
	CauseOfBlockage  string        `json:"causeOfBlockage,omitempty" description:"the cause of blockage"`
	Causes           []struct {
		Class            string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		ShortDescription string `json:"shortDescription,omitempty" description:"short description"`
		UserID           string `json:"userId,omitempty" description:"user id"`
		UserName         string `json:"userName,omitempty" description:"user name"`
	} `json:"causes,omitempty"`
	ChangeSet                 []interface{} `json:"changeSet,omitempty" description:"changeset information"`
	Description               interface{}   `json:"description,omitempty" description:"description"`
	DurationInMillis          interface{}   `json:"durationInMillis,omitempty" description:"duration time in millis"`
	EnQueueTime               interface{}   `json:"enQueueTime,omitempty" description:"the time of enter the queue"`
	EndTime                   interface{}   `json:"endTime,omitempty" description:"the time of end"`
	EstimatedDurationInMillis interface{}   `json:"estimatedDurationInMillis,omitempty" description:"estimated duration time, unit is millis"`
	ID                        string        `json:"id,omitempty" description:"id"`
	Name                      interface{}   `json:"name,omitempty" description:"name"`
	Organization              string        `json:"organization,omitempty" description:"the name of organization"`
	Pipeline                  string        `json:"pipeline,omitempty" description:"pipeline"`
	Replayable                bool          `json:"replayable,omitempty" description:"replayable or not"`
	Result                    string        `json:"result,omitempty" description:"the result of pipeline run. e.g. SUCCESS"`
	RunSummary                interface{}   `json:"runSummary,omitempty" description:"pipeline run summary"`
	StartTime                 interface{}   `json:"startTime,omitempty" description:"the time of start"`
	State                     string        `json:"state,omitempty" description:"run state. e.g. RUNNING"`
	Type                      string        `json:"type,omitempty" description:"type"`
	QueueID                   string        `json:"queueId,omitempty" description:"queue id"`
}

// GetArtifacts
type Artifacts struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Downloadable bool   `json:"downloadable,omitempty" description:"downloadable or not"`
	ID           string `json:"id,omitempty" description:"id"`
	Name         string `json:"name,omitempty" description:"name"`
	Path         string `json:"path,omitempty" description:"path"`
	Size         int    `json:"size,omitempty" description:"size"`
	URL          string `json:"url,omitempty" description:"The url for Download artifacts"`
}

// GetPipeBranch
type PipeBranch struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		Scm struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"scm,omitempty"`
		Actions struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"actions,omitempty"`
		Runs struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"runs,omitempty"`
		Trends struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"trends,omitempty"`
		Queue struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"queue,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Actions                   []interface{} `json:"actions,omitempty" description:"the list of all actions."`
	Disabled                  bool          `json:"disabled,omitempty" description:"disable or not, if disabled, can not do any action"`
	DisplayName               string        `json:"displayName,omitempty" description:"display name"`
	EstimatedDurationInMillis int           `json:"estimatedDurationInMillis,omitempty" description:"estimated duration time, unit is millis"`
	FullDisplayName           string        `json:"fullDisplayName,omitempty" description:"full display name"`
	FullName                  string        `json:"fullName,omitempty" description:"full name"`
	LatestRun                 struct {
		Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		Links struct {
			PrevRun struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"prevRun,omitempty"`
			Parent struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"parent,omitempty"`
			Tests struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"tests,omitempty"`
			Log struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"log,omitempty"`
			Self struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"self,omitempty"`
			BlueTestSummary struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"blueTestSummary,omitempty"`
			Actions struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"actions,omitempty"`
			Artifacts struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"artifacts,omitempty"`
		} `json:"_links,omitempty" description:"references the reachable path to this resource"`
		Actions          []interface{} `json:"actions,omitempty" description:"the list of all actions"`
		ArtifactsZipFile string        `json:"artifactsZipFile,omitempty" description:"the artifacts zip file"`
		CauseOfBlockage  interface{}   `json:"causeOfBlockage,omitempty" description:"the cause of blockage"`
		Causes           []struct {
			Class            string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
			ShortDescription string `json:"shortDescription,omitempty" description:"short description"`
		} `json:"causes,omitempty"`
		ChangeSet                 []interface{} `json:"changeSet,omitempty" description:"changeset information"`
		Description               interface{}   `json:"description,omitempty" description:"description"`
		DurationInMillis          int           `json:"durationInMillis,omitempty" description:"duration time in millis"`
		EnQueueTime               string        `json:"enQueueTime,omitempty" description:"the time of enter the queue"`
		EndTime                   string        `json:"endTime,omitempty" description:"the time of end"`
		EstimatedDurationInMillis int           `json:"estimatedDurationInMillis,omitempty" description:"estimated duration time in millis"`
		ID                        string        `json:"id,omitempty" description:"id"`
		Name                      interface{}   `json:"name,omitempty" description:"name"`
		Organization              string        `json:"organization,omitempty" description:"the name of organization"`
		Pipeline                  string        `json:"pipeline,omitempty" description:"pipeline"`
		Replayable                bool          `json:"replayable,omitempty" description:"replayable or not"`
		Result                    string        `json:"result,omitempty" description:"the result of pipeline run. e.g. SUCCESS"`
		RunSummary                string        `json:"runSummary,omitempty" description:"pipeline run summary"`
		StartTime                 string        `json:"startTime,omitempty" description:"start run"`
		State                     string        `json:"state,omitempty" description:"run state. e.g. RUNNING"`
		Type                      string        `json:"type,omitempty" description:"type"`
	} `json:"latestRun,omitempty"`
	Name         string `json:"name,omitempty" description:"name"`
	Organization string `json:"organization,omitempty" description:"the name of organization"`
	Parameters   []struct {
		Class                 string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		DefaultParameterValue struct {
			Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
			Name  string `json:"name,omitempty" description:"name"`
			Value string `json:"value,omitempty" description:"value"`
		} `json:"defaultParameterValue,omitempty"`
		Description string `json:"description,omitempty" description:"description"`
		Name        string `json:"name,omitempty" description:"name"`
		Type        string `json:"type,omitempty" description:"type"`
	} `json:"parameters,omitempty"`
	Permissions struct {
		Create    bool `json:"create,omitempty" description:"create action"`
		Configure bool `json:"configure,omitempty" description:"configure action"`
		Read      bool `json:"read,omitempty" description:"read action"`
		Start     bool `json:"start,omitempty" description:"start action"`
		Stop      bool `json:"stop,omitempty" description:"stop action"`
	} `json:"permissions,omitempty"`
	WeatherScore int `json:"weatherScore,omitempty" description:"the score to description the result of pipeline"`
	Branch       struct {
		IsPrimary bool          `json:"isPrimary,omitempty" description:"primary or not"`
		Issues    []interface{} `json:"issues,omitempty" description:"issues"`
		URL       string        `json:"url,omitempty" description:"url"`
	} `json:"branch,omitempty"`
}

// RunPipeline
type RunPayload struct {
	Parameters []struct {
		Name  string `json:"name,omitempty" description:"name"`
		Value string `json:"value,omitempty" description:"value"`
	} `json:"parameters,omitempty"`
}

type QueuedBlueRun struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		Parent struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"parent,omitempty"`
		Tests struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"tests,omitempty"`
		Log struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"log,omitempty"`
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		BlueTestSummary struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"blueTestSummary,omitempty"`
		Actions struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"actions,omitempty"`
		Artifacts struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"artifacts,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Actions          []interface{} `json:"actions,omitempty" description:"the list of all actions"`
	ArtifactsZipFile interface{}   `json:"artifactsZipFile,omitempty" description:"the artifacts zip file"`
	CauseOfBlockage  string        `json:"causeOfBlockage,omitempty" description:"the cause of blockage"`
	Causes           []struct {
		Class            string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		ShortDescription string `json:"shortDescription,omitempty" description:"short description"`
		UserID           string `json:"userId,omitempty" description:"user id"`
		UserName         string `json:"userName,omitempty" description:"user name"`
	} `json:"causes,omitempty"`
	ChangeSet                 []interface{} `json:"changeSet,omitempty" description:"changeset information"`
	Description               interface{}   `json:"description,omitempty" description:"description"`
	DurationInMillis          interface{}   `json:"durationInMillis,omitempty" description:"duration time in millis"`
	EnQueueTime               interface{}   `json:"enQueueTime,omitempty" description:"the time of enter the queue"`
	EndTime                   interface{}   `json:"endTime,omitempty" description:"the time of end"`
	EstimatedDurationInMillis interface{}   `json:"estimatedDurationInMillis,omitempty" description:"estimated duration time in millis"`
	ID                        string        `json:"id,omitempty" description:"id"`
	Name                      interface{}   `json:"name,omitempty" description:"name"`
	Organization              string        `json:"organization,omitempty" description:"the name of organization"`
	Pipeline                  string        `json:"pipeline,omitempty" description:"pipeline"`
	Replayable                bool          `json:"replayable,omitempty" description:"replayable or not"`
	Result                    string        `json:"result,omitempty" description:"the result of pipeline run. e.g. SUCCESS"`
	RunSummary                interface{}   `json:"runSummary,omitempty" description:"pipeline run summary"`
	StartTime                 interface{}   `json:"startTime,omitempty" description:"the time of start"`
	State                     string        `json:"state,omitempty" description:"run state. e.g. RUNNING"`
	Type                      string        `json:"type,omitempty" description:"type"`
	QueueID                   string        `json:"queueId,omitempty" description:"queue id"`
}

// GetNodeStatus
type NodeStatus struct {
	Class string `json:"_class,omitempty" description:""`
	Links struct {
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		Actions struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"actions,omitempty"`
		Steps struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"steps,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Actions            []interface{} `json:"actions,omitempty" description:"the list of all actions"`
	DisplayDescription interface{}   `json:"displayDescription,omitempty" description:"display description"`
	DisplayName        string        `json:"displayName,omitempty" description:"display name"`
	DurationInMillis   int           `json:"durationInMillis,omitempty" description:"duration time in millis"`
	ID                 string        `json:"id,omitempty" description:"id"`
	Input              *Input        `json:"input,omitempty" description:"the action should user input"`
	Result             string        `json:"result,omitempty" description:"the result of pipeline run. e.g. SUCCESS"`
	StartTime          string        `json:"startTime,omitempty" description:"the time of start"`
	State              string        `json:"state,omitempty" description:"run state. e.g. RUNNING"`
	Type               string        `json:"type,omitempty" description:"type"`
	CauseOfBlockage    interface{}   `json:"causeOfBlockage,omitempty" description:"the cause of blockage"`
	Edges              []struct {
		Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		ID    string `json:"id,omitempty" description:"id"`
		Type  string `json:"type,omitempty" description:"type"`
	} `json:"edges,omitempty"`
	FirstParent interface{} `json:"firstParent,omitempty" description:"first parent"`
	Restartable bool        `json:"restartable,omitempty" description:"restartable or not"`
	Steps       []struct {
		Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		Links struct {
			Self struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"self,omitempty"`
			Actions struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"actions,omitempty"`
		} `json:"_links,omitempty" description:"references the reachable path to this resource"`
		Actions []struct {
			Class string `json:"_class,omitempty" description:"references the reachable path to this resource"`
			Links struct {
				Self struct {
					Class string `json:"_class,omitempty"`
					Href  string `json:"href,omitempty"`
				} `json:"self,omitempty" description:""`
			} `json:"_links,omitempty" description:"references the reachable path to this resource"`
			URLName string `json:"urlName,omitempty" description:"url name"`
		} `json:"actions,omitempty"`
		DisplayDescription interface{} `json:"displayDescription,omitempty" description:"display description"`
		DisplayName        string      `json:"displayName,omitempty" description:"display name"`
		DurationInMillis   int         `json:"durationInMillis,omitempty" description:"duration time in millis"`
		ID                 string      `json:"id,omitempty" description:"id"`
		Input              *Input      `json:"input,omitempty" description:"the action should user input"`
		Result             string      `json:"result,omitempty" description:"the result of pipeline run. e.g. SUCCESS"`
		StartTime          string      `json:"startTime,omitempty" description:"the time of start"`
		State              string      `json:"state,omitempty" description:"run state. e.g. RUNNING"`
		Type               string      `json:"type,omitempty" description:"type"`
	} `json:"steps,omitempty"`
}

// CheckPipeline
type CheckPlayload struct {
	ID         string                    `json:"id,omitempty" description:"id"`
	Parameters []CheckPlayloadParameters `json:"parameters,omitempty"`
	Abort      bool                      `json:"abort,omitempty" description:"abort or not"`
}

type CreateScmServerReq struct {
	Name   string `json:"name,omitempty" description:"name of scm server"`
	ApiURL string `json:"apiUrl,omitempty"  description:"url of scm server"`
}

type CheckPlayloadParameters struct {
	Name  string      `json:"name,omitempty" description:"name"`
	Value interface{} `json:"value,omitempty" description:"value"`
}

// Getcrumb
type Crumb struct {
	Class             string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Crumb             string `json:"crumb,omitempty" description:"crumb data"`
	CrumbRequestField string `json:"crumbRequestField,omitempty" description:"crumb request field"`
}

// CheckScriptCompile
type CheckScript struct {
	Column  int    `json:"column,omitempty" description:"column e.g. 0"`
	Line    int    `json:"line,omitempty" description:"line e.g. 0"`
	Message string `json:"message,omitempty" description:"message e.g. unexpected char: '#'"`
	Status  string `json:"status,omitempty" description:"status e.g. fail"`
}

// CheckCron
type CronData struct {
	PipelineName string `json:"pipelineName,omitempty" description:"Pipeline name, if pipeline haven't created, not required'"`
	Cron         string `json:"cron" description:"Cron script data."`
}

type CheckCronRes struct {
	Result   string `json:"result,omitempty" description:"result e.g. ok, error"`
	Message  string `json:"message,omitempty" description:"message"`
	LastTime string `json:"lastTime,omitempty" description:"last run time."`
	NextTime string `json:"nextTime,omitempty" description:"next run time."`
}

// GetPipelineRun
type PipelineRun struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		PrevRun struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"prevRun,omitempty"`
		Parent struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"parent,omitempty"`
		Tests struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"tests,omitempty"`
		Nodes struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"nodes,omitempty"`
		Log struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"log,omitempty"`
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		BlueTestSummary struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"blueTestSummary,omitempty"`
		Actions struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"actions,omitempty"`
		Steps struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"steps,omitempty"`
		Artifacts struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"artifacts,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Actions          []interface{} `json:"actions,omitempty" description:"the list of all actions"`
	ArtifactsZipFile interface{}   `json:"artifactsZipFile,omitempty" description:"the artifacts zip file"`
	CauseOfBlockage  interface{}   `json:"causeOfBlockage,omitempty" description:"the cause of blockage"`
	Causes           []struct {
		Class            string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		ShortDescription string `json:"shortDescription,omitempty" description:"short description"`
		UserID           string `json:"userId,omitempty" description:"user id"`
		UserName         string `json:"userName,omitempty" description:"user name"`
	} `json:"causes,omitempty"`
	ChangeSet                 []interface{} `json:"changeSet,omitempty" description:"changeset information"`
	Description               interface{}   `json:"description,omitempty" description:"description"`
	DurationInMillis          int           `json:"durationInMillis,omitempty" description:"duration time in millis"`
	EnQueueTime               string        `json:"enQueueTime,omitempty" description:"the time of enter the queue"`
	EndTime                   string        `json:"endTime,omitempty" description:"the time of end"`
	EstimatedDurationInMillis int           `json:"estimatedDurationInMillis,omitempty" description:"estimated duration time in millis"`
	ID                        string        `json:"id,omitempty" description:"id"`
	Name                      interface{}   `json:"name,omitempty" description:"name"`
	Organization              string        `json:"organization,omitempty" description:"the name of organization"`
	Pipeline                  string        `json:"pipeline,omitempty" description:"the name of pipeline"`
	Replayable                bool          `json:"replayable,omitempty" description:"replayable or not"`
	Result                    string        `json:"result,omitempty" description:"the result of pipeline run. e.g. SUCCESS"`
	RunSummary                string        `json:"runSummary,omitempty" description:"pipeline run summary"`
	StartTime                 string        `json:"startTime,omitempty" description:"the time of start"`
	State                     string        `json:"state,omitempty" description:"run state. e.g. RUNNING"`
	Type                      string        `json:"type,omitempty" description:"type"`
	Branch                    interface{}   `json:"branch,omitempty" description:"branch"`
	CommitID                  interface{}   `json:"commitId,omitempty" description:"commit id"`
	CommitURL                 interface{}   `json:"commitUrl,omitempty" description:"commit url"`
	PullRequest               interface{}   `json:"pullRequest,omitempty" description:"pull request"`
}

// GetBranchPipeRun
type BranchPipeline struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		Scm struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"scm,omitempty"`
		Actions struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"actions,omitempty"`
		Runs struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"runs,omitempty"`
		Trends struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"trends,omitempty"`
		Queue struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"queue,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Actions                   []interface{} `json:"actions,omitempty" description:"the list of all actions"`
	Disabled                  bool          `json:"disabled,omitempty" description:"disable or not, if disabled, can not do any action"`
	DisplayName               string        `json:"displayName,omitempty" description:"display name"`
	EstimatedDurationInMillis int           `json:"estimatedDurationInMillis,omitempty" description:"estimated duration time in millis"`
	FullDisplayName           string        `json:"fullDisplayName,omitempty" description:"full display name"`
	FullName                  string        `json:"fullName,omitempty" description:"full name"`
	LatestRun                 struct {
		Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		Links struct {
			PrevRun struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"prevRun,omitempty"`
			Parent struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"parent,omitempty"`
			Tests struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"tests,omitempty"`
			Log struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"log,omitempty"`
			Self struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"self,omitempty"`
			BlueTestSummary struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"blueTestSummary,omitempty"`
			Actions struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"actions,omitempty"`
			Artifacts struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"artifacts,omitempty"`
		} `json:"_links,omitempty" description:"references the reachable path to this resource"`
		Actions          []interface{} `json:"actions,omitempty" description:"the list of all actions"`
		ArtifactsZipFile string        `json:"artifactsZipFile,omitempty" description:"the artifacts zip file"`
		CauseOfBlockage  interface{}   `json:"causeOfBlockage,omitempty" description:"the cause of blockage"`
		Causes           []struct {
			Class            string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
			ShortDescription string `json:"shortDescription,omitempty" description:"short description"`
			UserID           string `json:"userId,omitempty" description:"user id"`
			UserName         string `json:"userName,omitempty" description:"user name"`
		} `json:"causes,omitempty"`
		ChangeSet                 []interface{} `json:"changeSet,omitempty" description:"changeset information"`
		Description               interface{}   `json:"description,omitempty" description:"description"`
		DurationInMillis          int           `json:"durationInMillis,omitempty" description:"duration time in millis"`
		EnQueueTime               string        `json:"enQueueTime,omitempty" description:"the time of enter the queue"`
		EndTime                   string        `json:"endTime,omitempty" description:"the time of end"`
		EstimatedDurationInMillis int           `json:"estimatedDurationInMillis,omitempty" description:"estimated duration time in millis"`
		ID                        string        `json:"id,omitempty" description:"id"`
		Name                      interface{}   `json:"name,omitempty" description:"name"`
		Organization              string        `json:"organization,omitempty" description:"the name of organization"`
		Pipeline                  string        `json:"pipeline,omitempty" description:"pipeline"`
		Replayable                bool          `json:"replayable,omitempty" description:"Replayable or not"`
		Result                    string        `json:"result,omitempty" description:"the result of pipeline run. e.g. SUCCESS"`
		RunSummary                string        `json:"runSummary,omitempty" description:"pipeline run summary"`
		StartTime                 string        `json:"startTime,omitempty" description:"the time of start"`
		State                     string        `json:"state,omitempty" description:"run state. e.g. RUNNING"`
		Type                      string        `json:"type,omitempty" description:"type"`
	} `json:"latestRun,omitempty"`
	Name         string `json:"name,omitempty" description:"name"`
	Organization string `json:"organization,omitempty" description:"the name of organization"`
	Parameters   []struct {
		Class                 string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		DefaultParameterValue struct {
			Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
			Name  string `json:"name,omitempty" description:"name"`
			Value string `json:"value,omitempty" description:"value"`
		} `json:"defaultParameterValue,omitempty" description:""`
		Description string `json:"description,omitempty" description:"description"`
		Name        string `json:"name,omitempty" description:"name"`
		Type        string `json:"type,omitempty" description:"type"`
	} `json:"parameters,omitempty"`
	Permissions struct {
		Create    bool `json:"create,omitempty" description:"create action"`
		Configure bool `json:"configure,omitempty" description:"configure action"`
		Read      bool `json:"read,omitempty" description:"read action"`
		Start     bool `json:"start,omitempty" description:"start action"`
		Stop      bool `json:"stop,omitempty" description:"stop action"`
	} `json:"permissions,omitempty"`
	WeatherScore int `json:"weatherScore,omitempty" description:"the score to description the result of pipeline"`
	Branch       struct {
		IsPrimary bool          `json:"isPrimary,omitempty" description:"primary or not"`
		Issues    []interface{} `json:"issues,omitempty" description:"issues"`
		URL       string        `json:"url,omitempty" description:"url"`
	} `json:"branch,omitempty"`
}

// GetPipelineRunNodes
type PipelineRunNodes struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		Actions struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"actions,omitempty"`
		Steps struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"steps,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Actions            []interface{} `json:"actions,omitempty" description:"the list of all actions"`
	DisplayDescription interface{}   `json:"displayDescription,omitempty" description:"display description"`
	DisplayName        string        `json:"displayName,omitempty" description:"display name"`
	DurationInMillis   int           `json:"durationInMillis,omitempty" description:"duration time in mullis"`
	ID                 string        `json:"id,omitempty" description:"id"`
	Input              *Input        `json:"input,omitempty" description:"the action should user input"`
	Result             string        `json:"result,omitempty" description:"the result of pipeline run. e.g. SUCCESS"`
	StartTime          string        `json:"startTime,omitempty" description:"the time of start"`
	State              string        `json:"state,omitempty" description:"run state. e.g. FINISHED"`
	Type               string        `json:"type,omitempty" description:"type"`
	CauseOfBlockage    interface{}   `json:"causeOfBlockage,omitempty" description:"the cause of blockage"`
	Edges              []interface{} `json:"edges,omitempty" description:"edges"`
	FirstParent        interface{}   `json:"firstParent,omitempty" description:"first parent"`
	Restartable        bool          `json:"restartable,omitempty" description:"restartable or not"`
}

// GetNodeSteps
type NodeSteps struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		Actions struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"actions,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Actions []struct {
		Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		Links struct {
			Self struct {
				Class string `json:"_class,omitempty"`
				Href  string `json:"href,omitempty"`
			} `json:"self,omitempty"`
		} `json:"_links,omitempty" description:"references the reachable path to this resource"`
		URLName string `json:"urlName,omitempty" description:"url name"`
	} `json:"actions,omitempty"`
	DisplayDescription string `json:"displayDescription,omitempty" description:"display description"`
	DisplayName        string `json:"displayName,omitempty" description:"display name"`
	DurationInMillis   int    `json:"durationInMillis,omitempty" description:"duration time in mullis"`
	ID                 string `json:"id,omitempty" description:"id"`
	Input              *Input `json:"input,omitempty" description:"the action should user input"`
	Result             string `json:"result,omitempty" description:"the result of pipeline run. e.g. SUCCESS"`
	StartTime          string `json:"startTime,omitempty" description:"the time of starts"`
	State              string `json:"state,omitempty" description:"run state. e.g. SKIPPED"`
	Type               string `json:"type,omitempty" description:"type"`
}

// CheckScriptCompile
type ReqScript struct {
	Value string `json:"value,omitempty" description:"Pipeline script data"`
}

// ToJenkinsfile requests
type ReqJson struct {
	Json string `json:"json,omitempty" description:"json data"`
}

// ToJenkinsfile response
type ResJenkinsfile struct {
	Status string `json:"status,omitempty" description:"status e.g. ok"`
	Data   struct {
		Result      string `json:"result,omitempty" description:"result e.g. success"`
		Jenkinsfile string `json:"jenkinsfile,omitempty" description:"jenkinsfile"`
		Errors      []struct {
			Location []string `json:"location,omitempty" description:"err location"`
			Error    string   `json:"error,omitempty" description:"error message"`
		} `json:"errors,omitempty"`
	} `json:"data,omitempty"`
}

type ReqJenkinsfile struct {
	Jenkinsfile string `json:"jenkinsfile,omitempty" description:"jenkinsfile"`
}

type ResJson struct {
	Status string `json:"status,omitempty" description:"status e.g. ok"`
	Data   struct {
		Result string `json:"result,omitempty" description:"result e.g. success"`
		JSON   struct {
			Pipeline struct {
				Stages []interface{} `json:"stages,omitempty" description:"stages"`
				Agent  struct {
					Type      string `json:"type,omitempty" description:"type"`
					Arguments []struct {
						Key   string `json:"key,omitempty" description:"key"`
						Value struct {
							IsLiteral bool   `json:"isLiteral,omitempty" description:"is literal or not"`
							Value     string `json:"value,omitempty" description:"value"`
						} `json:"value,omitempty"`
					} `json:"arguments,omitempty"`
				} `json:"agent,omitempty"`
			} `json:"pipeline,omitempty"`
		} `json:"json,omitempty"`
	} `json:"data,omitempty"`
}

type NodesDetail struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links struct {
		Self struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		Actions struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"actions,omitempty"`
		Steps struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"steps,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	Actions            []interface{} `json:"actions,omitempty" description:"the list of all actions"`
	DisplayDescription interface{}   `json:"displayDescription,omitempty" description:"display description"`
	DisplayName        string        `json:"displayName,omitempty" description:"display name"`
	DurationInMillis   int           `json:"durationInMillis,omitempty" description:"duration time in millis"`
	ID                 string        `json:"id,omitempty" description:"id"`
	Input              *Input        `json:"input,omitempty" description:"the action should user input"`
	Result             string        `json:"result,omitempty" description:"the result of pipeline run. e.g. SUCCESS"`
	StartTime          string        `json:"startTime,omitempty" description:"the time of start"`
	State              string        `json:"state,omitempty" description:"run state. e.g. SKIPPED"`
	Type               string        `json:"type,omitempty" description:"type"`
	CauseOfBlockage    interface{}   `json:"causeOfBlockage,omitempty" description:"the cause of blockage"`
	Edges              []struct {
		Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
		ID    string `json:"id,omitempty" description:"id"`
		Type  string `json:"type,omitempty" description:"type"`
	} `json:"edges,omitempty"`
	FirstParent interface{} `json:"firstParent,omitempty" description:"first parent"`
	Restartable bool        `json:"restartable,omitempty" description:"restartable or not"`
	Steps       []NodeSteps `json:"steps,omitempty" description:"steps"`
}

type NodesStepsIndex struct {
	Id    int         `json:"id,omitempty" description:"id"`
	Steps []NodeSteps `json:"steps,omitempty" description:"steps"`
}

type Input struct {
	Class string `json:"_class,omitempty" description:"It’s a fully qualified name and is an identifier of the producer of this resource's capability."`
	Links *struct {
		Self *struct {
			Class string `json:"_class,omitempty"`
			Href  string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty" description:"references the reachable path to this resource"`
	ID         string        `json:"id,omitempty" description:"the id of check action"`
	Message    string        `json:"message,omitempty" description:"the message of check action"`
	Ok         string        `json:"ok,omitempty" description:"check status. e.g. \"Proceed\""`
	Parameters []interface{} `json:"parameters,omitempty" description:"the parameters of check action"`
	Submitter  interface{}   `json:"submitter,omitempty" description:"check submitter"`
}

type PipelineOperator interface {
	GetPipeline(projectName, pipelineName string, req *http.Request) ([]byte, error)

	ListPipelines(req *http.Request) ([]byte, error)

	GetPipelineRun(projectName, pipelineName, runId string, req *http.Request) ([]byte, error)

	ListPipelineRuns(projectName, pipelineName string, req *http.Request) ([]byte, error)

	StopPipeline(projectName, pipelineName, runId string, req *http.Request) ([]byte, error)

	ReplayPipeline(projectName, pipelineName, runId string, req *http.Request) ([]byte, error)

	RunPipeline(projectName, pipelineName string, req *http.Request) ([]byte, error)
}
const (
	NoScmPipelineType       = "pipeline"
	MultiBranchPipelineType = "multi-branch-pipeline"
)

type Parameters []*Parameter

var ParameterTypeMap = map[string]string{
	"hudson.model.StringParameterDefinition":   "string",
	"hudson.model.ChoiceParameterDefinition":   "choice",
	"hudson.model.TextParameterDefinition":     "text",
	"hudson.model.BooleanParameterDefinition":  "boolean",
	"hudson.model.FileParameterDefinition":     "file",
	"hudson.model.PasswordParameterDefinition": "password",
}

type ProjectPipeline struct {
	Type                string               `json:"type" description:"type of devops pipeline, in scm or no scm"`
	Pipeline            *NoScmPipeline       `json:"pipeline,omitempty" description:"no scm pipeline structs"`
	MultiBranchPipeline *MultiBranchPipeline `json:"multi_branch_pipeline,omitempty" description:"in scm pipeline structs"`
}

type NoScmPipeline struct {
	Name              string             `json:"name" description:"name of pipeline"`
	Description       string             `json:"descriptio,omitempty" description:"description of pipeline"`
	Discarder         *DiscarderProperty `json:"discarder,omitempty" description:"Discarder of pipeline, managing when to drop a pipeline"`
	Parameters        *Parameters        `json:"parameters,omitempty" description:"Parameters define of pipeline,user could pass param when run pipeline"`
	DisableConcurrent bool               `json:"disable_concurrent,omitempty" mapstructure:"disable_concurrent" description:"Whether to prohibit the pipeline from running in parallel"`
	TimerTrigger      *TimerTrigger      `json:"timer_trigger,omitempty" mapstructure:"timer_trigger" description:"Timer to trigger pipeline run"`
	RemoteTrigger     *RemoteTrigger     `json:"remote_trigger,omitempty" mapstructure:"remote_trigger" description:"Remote api define to trigger pipeline run"`
	Jenkinsfile       string             `json:"jenkinsfile,omitempty" description:"Jenkinsfile's content'"`
}

type MultiBranchPipeline struct {
	Name                  string                 `json:"name" description:"name of pipeline"`
	Description           string                 `json:"descriptio,omitempty" description:"description of pipeline"`
	Discarder             *DiscarderProperty     `json:"discarder,omitempty" description:"Discarder of pipeline, managing when to drop a pipeline"`
	TimerTrigger          *TimerTrigger          `json:"timer_trigger,omitempty" mapstructure:"timer_trigger" description:"Timer to trigger pipeline run"`
	SourceType            string                 `json:"source_type" description:"type of scm, such as github/git/svn"`
	GitSource             *GitSource             `json:"git_source,omitempty" description:"git scm define"`
	GitHubSource          *GithubSource          `json:"github_source,omitempty" description:"github scm define"`
	SvnSource             *SvnSource             `json:"svn_source,omitempty" description:"multi branch svn scm define"`
	SingleSvnSource       *SingleSvnSource       `json:"single_svn_source,omitempty" description:"single branch svn scm define"`
	BitbucketServerSource *BitbucketServerSource `json:"bitbucket_server_source,omitempty" description:"bitbucket server scm defile"`
	ScriptPath            string                 `json:"script_path" mapstructure:"script_path" description:"script path in scm"`
	MultiBranchJobTrigger *MultiBranchJobTrigger `json:"multibranch_job_trigger,omitempty" mapstructure:"multibranch_job_trigger" description:"Pipeline tasks that need to be triggered when branch creation/deletion"`
}

type GitSource struct {
	ScmId            string          `json:"scm_id,omitempty" description:"uid of scm"`
	Url              string          `json:"url,omitempty" mapstructure:"url" description:"url of git source"`
	CredentialId     string          `json:"credential_id,omitempty" mapstructure:"credential_id" description:"credential id to access git source"`
	DiscoverBranches bool            `json:"discover_branches,omitempty" mapstructure:"discover_branches" description:"Whether to discover a branch"`
	CloneOption      *GitCloneOption `json:"git_clone_option,omitempty" mapstructure:"git_clone_option" description:"advavced git clone options"`
	RegexFilter      string          `json:"regex_filter,omitempty" mapstructure:"regex_filter" description:"Regex used to match the name of the branch that needs to be run"`
}

type GithubSource struct {
	ScmId                string               `json:"scm_id,omitempty" description:"uid of scm"`
	Owner                string               `json:"owner,omitempty" mapstructure:"owner" description:"owner of github repo"`
	Repo                 string               `json:"repo,omitempty" mapstructure:"repo" description:"repo name of github repo"`
	CredentialId         string               `json:"credential_id,omitempty" mapstructure:"credential_id" description:"credential id to access github source"`
	ApiUri               string               `json:"api_uri,omitempty" mapstructure:"api_uri" description:"The api url can specify the location of the github apiserver.For private cloud configuration"`
	DiscoverBranches     int                  `json:"discover_branches,omitempty" mapstructure:"discover_branches" description:"Discover branch configuration"`
	DiscoverPRFromOrigin int                  `json:"discover_pr_from_origin,omitempty" mapstructure:"discover_pr_from_origin" description:"Discover origin PR configuration"`
	DiscoverPRFromForks  *DiscoverPRFromForks `json:"discover_pr_from_forks,omitempty" mapstructure:"discover_pr_from_forks" description:"Discover fork PR configuration"`
	CloneOption          *GitCloneOption      `json:"git_clone_option,omitempty" mapstructure:"git_clone_option" description:"advavced git clone options"`
	RegexFilter          string               `json:"regex_filter,omitempty" mapstructure:"regex_filter" description:"Regex used to match the name of the branch that needs to be run"`
}

type MultiBranchJobTrigger struct {
	CreateActionJobsToTrigger string `json:"create_action_job_to_trigger,omitempty" description:"pipeline name to trigger"`
	DeleteActionJobsToTrigger string `json:"delete_action_job_to_trigger,omitempty" description:"pipeline name to trigger"`
}

type BitbucketServerSource struct {
	ScmId                string               `json:"scm_id,omitempty" description:"uid of scm"`
	Owner                string               `json:"owner,omitempty" mapstructure:"owner" description:"owner of github repo"`
	Repo                 string               `json:"repo,omitempty" mapstructure:"repo" description:"repo name of github repo"`
	CredentialId         string               `json:"credential_id,omitempty" mapstructure:"credential_id" description:"credential id to access github source"`
	ApiUri               string               `json:"api_uri,omitempty" mapstructure:"api_uri" description:"The api url can specify the location of the github apiserver.For private cloud configuration"`
	DiscoverBranches     int                  `json:"discover_branches,omitempty" mapstructure:"discover_branches" description:"Discover branch configuration"`
	DiscoverPRFromOrigin int                  `json:"discover_pr_from_origin,omitempty" mapstructure:"discover_pr_from_origin" description:"Discover origin PR configuration"`
	DiscoverPRFromForks  *DiscoverPRFromForks `json:"discover_pr_from_forks,omitempty" mapstructure:"discover_pr_from_forks" description:"Discover fork PR configuration"`
	CloneOption          *GitCloneOption      `json:"git_clone_option,omitempty" mapstructure:"git_clone_option" description:"advavced git clone options"`
	RegexFilter          string               `json:"regex_filter,omitempty" mapstructure:"regex_filter" description:"Regex used to match the name of the branch that needs to be run"`
}

type GitCloneOption struct {
	Shallow bool `json:"shallow,omitempty" mapstructure:"shallow" description:"Whether to use git shallow clone"`
	Timeout int  `json:"timeout,omitempty" mapstructure:"timeout" description:"git clone timeout mins"`
	Depth   int  `json:"depth,omitempty" mapstructure:"depth" description:"git clone depth"`
}

type SvnSource struct {
	ScmId        string `json:"scm_id,omitempty" description:"uid of scm"`
	Remote       string `json:"remote,omitempty" description:"remote address url"`
	CredentialId string `json:"credential_id,omitempty" mapstructure:"credential_id" description:"credential id to access svn source"`
	Includes     string `json:"includes,omitempty" description:"branches to run pipeline"`
	Excludes     string `json:"excludes,omitempty" description:"branches do not run pipeline"`
}
type SingleSvnSource struct {
	ScmId        string `json:"scm_id,omitempty" description:"uid of scm"`
	Remote       string `json:"remote,omitempty" description:"remote address url"`
	CredentialId string `json:"credential_id,omitempty" mapstructure:"credential_id" description:"credential id to access svn source"`
}

type DiscoverPRFromForks struct {
	Strategy int `json:"strategy,omitempty" mapstructure:"strategy" description:"github discover strategy"`
	Trust    int `json:"trust,omitempty" mapstructure:"trust" description:"trust user type"`
}

type DiscarderProperty struct {
	DaysToKeep string `json:"days_to_keep,omitempty" mapstructure:"days_to_keep" description:"days to keep pipeline"`
	NumToKeep  string `json:"num_to_keep,omitempty" mapstructure:"num_to_keep" description:"nums to keep pipeline"`
}

type Parameter struct {
	Name         string `json:"name" description:"name of param"`
	DefaultValue string `json:"default_value,omitempty" mapstructure:"default_value" description:"default value of param"`
	Type         string `json:"type" description:"type of param"`
	Description  string `json:"description,omitempty" description:"description of pipeline"`
}

type TimerTrigger struct {
	// user in no scm job
	Cron string `json:"cron,omitempty" description:"jenkins cron script"`

	// use in multi-branch job
	Interval string `json:"interval,omitempty" description:"interval ms"`
}

type RemoteTrigger struct {
	Token string `json:"token,omitempty" description:"remote trigger token"`
}

type PipelineOperator interface {
	CreateProjectPipeline(projectId string, pipeline *ProjectPipeline) (string, error)
	DeleteProjectPipeline(projectId string, pipelineId string) (string, error)
	UpdateProjectPipeline(projectId string, pipeline *ProjectPipeline) (string, error)
	GetProjectPipelineConfig(projectId, pipelineId string) (*ProjectPipeline, error)
}
