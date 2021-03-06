package client

const (
	DATABASECHANGELOG_TYPE = "databasechangelog"
)

type Databasechangelog struct {
	Resource `yaml:"-"`

	Author string `json:"author,omitempty" yaml:"author,omitempty"`

	Comments string `json:"comments,omitempty" yaml:"comments,omitempty"`

	Contexts string `json:"contexts,omitempty" yaml:"contexts,omitempty"`

	Dateexecuted string `json:"dateexecuted,omitempty" yaml:"dateexecuted,omitempty"`

	DeploymentId string `json:"deploymentId,omitempty" yaml:"deployment_id,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Exectype string `json:"exectype,omitempty" yaml:"exectype,omitempty"`

	Filename string `json:"filename,omitempty" yaml:"filename,omitempty"`

	Labels string `json:"labels,omitempty" yaml:"labels,omitempty"`

	Liquibase string `json:"liquibase,omitempty" yaml:"liquibase,omitempty"`

	Md5sum string `json:"md5sum,omitempty" yaml:"md5sum,omitempty"`

	Orderexecuted int64 `json:"orderexecuted,omitempty" yaml:"orderexecuted,omitempty"`

	Tag string `json:"tag,omitempty" yaml:"tag,omitempty"`
}

type DatabasechangelogCollection struct {
	Collection
	Data   []Databasechangelog `json:"data,omitempty"`
	client *DatabasechangelogClient
}

type DatabasechangelogClient struct {
	rancherClient *RancherClient
}

type DatabasechangelogOperations interface {
	List(opts *ListOpts) (*DatabasechangelogCollection, error)
	Create(opts *Databasechangelog) (*Databasechangelog, error)
	Update(existing *Databasechangelog, updates interface{}) (*Databasechangelog, error)
	ById(id string) (*Databasechangelog, error)
	Delete(container *Databasechangelog) error
}

func newDatabasechangelogClient(rancherClient *RancherClient) *DatabasechangelogClient {
	return &DatabasechangelogClient{
		rancherClient: rancherClient,
	}
}

func (c *DatabasechangelogClient) Create(container *Databasechangelog) (*Databasechangelog, error) {
	resp := &Databasechangelog{}
	err := c.rancherClient.doCreate(DATABASECHANGELOG_TYPE, container, resp)
	return resp, err
}

func (c *DatabasechangelogClient) Update(existing *Databasechangelog, updates interface{}) (*Databasechangelog, error) {
	resp := &Databasechangelog{}
	err := c.rancherClient.doUpdate(DATABASECHANGELOG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DatabasechangelogClient) List(opts *ListOpts) (*DatabasechangelogCollection, error) {
	resp := &DatabasechangelogCollection{}
	err := c.rancherClient.doList(DATABASECHANGELOG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *DatabasechangelogCollection) Next() (*DatabasechangelogCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &DatabasechangelogCollection{}
		err := cc.client.rancherClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *DatabasechangelogClient) ById(id string) (*Databasechangelog, error) {
	resp := &Databasechangelog{}
	err := c.rancherClient.doById(DATABASECHANGELOG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *DatabasechangelogClient) Delete(container *Databasechangelog) error {
	return c.rancherClient.doResourceDelete(DATABASECHANGELOG_TYPE, &container.Resource)
}
