package client

const (
	ClusterLoggingSpecType                        = "clusterLoggingSpec"
	ClusterLoggingSpecFieldClusterID              = "clusterId"
	ClusterLoggingSpecFieldCustomTargetConfig     = "customTargetConfig"
	ClusterLoggingSpecFieldDisplayName            = "displayName"
	ClusterLoggingSpecFieldElasticsearchConfig    = "elasticsearchConfig"
	ClusterLoggingSpecFieldExcludeSystemComponent = "excludeSystemComponent"
	ClusterLoggingSpecFieldFluentForwarderConfig  = "fluentForwarderConfig"
	ClusterLoggingSpecFieldKafkaConfig            = "kafkaConfig"
	ClusterLoggingSpecFieldOutputFlushInterval    = "outputFlushInterval"
	ClusterLoggingSpecFieldOutputTags             = "outputTags"
	ClusterLoggingSpecFieldSplunkConfig           = "splunkConfig"
	ClusterLoggingSpecFieldSyslogConfig           = "syslogConfig"
)

type ClusterLoggingSpec struct {
	ClusterID              string                 `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	CustomTargetConfig     *CustomTargetConfig    `json:"customTargetConfig,omitempty" yaml:"customTargetConfig,omitempty"`
	DisplayName            string                 `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	ElasticsearchConfig    *ElasticsearchConfig   `json:"elasticsearchConfig,omitempty" yaml:"elasticsearchConfig,omitempty"`
	ExcludeSystemComponent bool                   `json:"excludeSystemComponent,omitempty" yaml:"excludeSystemComponent,omitempty"`
	FluentForwarderConfig  *FluentForwarderConfig `json:"fluentForwarderConfig,omitempty" yaml:"fluentForwarderConfig,omitempty"`
	KafkaConfig            *KafkaConfig           `json:"kafkaConfig,omitempty" yaml:"kafkaConfig,omitempty"`
	OutputFlushInterval    int64                  `json:"outputFlushInterval,omitempty" yaml:"outputFlushInterval,omitempty"`
	OutputTags             map[string]string      `json:"outputTags,omitempty" yaml:"outputTags,omitempty"`
	SplunkConfig           *SplunkConfig          `json:"splunkConfig,omitempty" yaml:"splunkConfig,omitempty"`
	SyslogConfig           *SyslogConfig          `json:"syslogConfig,omitempty" yaml:"syslogConfig,omitempty"`
}
