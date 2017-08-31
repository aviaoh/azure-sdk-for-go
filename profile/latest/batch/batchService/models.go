package batchservice

import (
	 original "github.com/Azure/azure-sdk-for-go/service/batch/2017-06-01.5.1/batchService"
)

type (
	 AccountClient = original.AccountClient
	 ApplicationClient = original.ApplicationClient
	 CertificateClient = original.CertificateClient
	 ComputeNodeClient = original.ComputeNodeClient
	 AccessScope = original.AccessScope
	 AllocationState = original.AllocationState
	 AutoUserScope = original.AutoUserScope
	 CachingType = original.CachingType
	 CertificateFormat = original.CertificateFormat
	 CertificateState = original.CertificateState
	 CertificateStoreLocation = original.CertificateStoreLocation
	 CertificateVisibility = original.CertificateVisibility
	 ComputeNodeDeallocationOption = original.ComputeNodeDeallocationOption
	 ComputeNodeFillType = original.ComputeNodeFillType
	 ComputeNodeRebootOption = original.ComputeNodeRebootOption
	 ComputeNodeReimageOption = original.ComputeNodeReimageOption
	 ComputeNodeState = original.ComputeNodeState
	 DependencyAction = original.DependencyAction
	 DisableComputeNodeSchedulingOption = original.DisableComputeNodeSchedulingOption
	 DisableJobOption = original.DisableJobOption
	 ElevationLevel = original.ElevationLevel
	 ErrorCategory = original.ErrorCategory
	 InboundEndpointProtocol = original.InboundEndpointProtocol
	 JobAction = original.JobAction
	 JobPreparationTaskState = original.JobPreparationTaskState
	 JobReleaseTaskState = original.JobReleaseTaskState
	 JobScheduleState = original.JobScheduleState
	 JobState = original.JobState
	 NetworkSecurityGroupRuleAccess = original.NetworkSecurityGroupRuleAccess
	 OnAllTasksComplete = original.OnAllTasksComplete
	 OnTaskFailure = original.OnTaskFailure
	 OSType = original.OSType
	 OutputFileUploadCondition = original.OutputFileUploadCondition
	 PoolLifetimeOption = original.PoolLifetimeOption
	 PoolState = original.PoolState
	 SchedulingState = original.SchedulingState
	 StartTaskState = original.StartTaskState
	 SubtaskState = original.SubtaskState
	 TaskAddStatus = original.TaskAddStatus
	 TaskCountValidationStatus = original.TaskCountValidationStatus
	 TaskExecutionResult = original.TaskExecutionResult
	 TaskState = original.TaskState
	 AccountListNodeAgentSkusResult = original.AccountListNodeAgentSkusResult
	 AffinityInformation = original.AffinityInformation
	 ApplicationListResult = original.ApplicationListResult
	 ApplicationPackageReference = original.ApplicationPackageReference
	 ApplicationSummary = original.ApplicationSummary
	 AuthenticationTokenSettings = original.AuthenticationTokenSettings
	 AutoPoolSpecification = original.AutoPoolSpecification
	 AutoScaleRun = original.AutoScaleRun
	 AutoScaleRunError = original.AutoScaleRunError
	 AutoUserSpecification = original.AutoUserSpecification
	 BatchError = original.BatchError
	 BatchErrorDetail = original.BatchErrorDetail
	 Certificate = original.Certificate
	 CertificateAddParameter = original.CertificateAddParameter
	 CertificateListResult = original.CertificateListResult
	 CertificateReference = original.CertificateReference
	 CloudJob = original.CloudJob
	 CloudJobListPreparationAndReleaseTaskStatusResult = original.CloudJobListPreparationAndReleaseTaskStatusResult
	 CloudJobListResult = original.CloudJobListResult
	 CloudJobSchedule = original.CloudJobSchedule
	 CloudJobScheduleListResult = original.CloudJobScheduleListResult
	 CloudPool = original.CloudPool
	 CloudPoolListResult = original.CloudPoolListResult
	 CloudServiceConfiguration = original.CloudServiceConfiguration
	 CloudTask = original.CloudTask
	 CloudTaskListResult = original.CloudTaskListResult
	 CloudTaskListSubtasksResult = original.CloudTaskListSubtasksResult
	 ComputeNode = original.ComputeNode
	 ComputeNodeEndpointConfiguration = original.ComputeNodeEndpointConfiguration
	 ComputeNodeError = original.ComputeNodeError
	 ComputeNodeGetRemoteLoginSettingsResult = original.ComputeNodeGetRemoteLoginSettingsResult
	 ComputeNodeInformation = original.ComputeNodeInformation
	 ComputeNodeListResult = original.ComputeNodeListResult
	 ComputeNodeUser = original.ComputeNodeUser
	 DeleteCertificateError = original.DeleteCertificateError
	 EnvironmentSetting = original.EnvironmentSetting
	 ErrorMessage = original.ErrorMessage
	 ExitCodeMapping = original.ExitCodeMapping
	 ExitCodeRangeMapping = original.ExitCodeRangeMapping
	 ExitConditions = original.ExitConditions
	 ExitOptions = original.ExitOptions
	 FileProperties = original.FileProperties
	 ImageReference = original.ImageReference
	 InboundEndpoint = original.InboundEndpoint
	 InboundNATPool = original.InboundNATPool
	 JobAddParameter = original.JobAddParameter
	 JobConstraints = original.JobConstraints
	 JobDisableParameter = original.JobDisableParameter
	 JobExecutionInformation = original.JobExecutionInformation
	 JobManagerTask = original.JobManagerTask
	 JobPatchParameter = original.JobPatchParameter
	 JobPreparationAndReleaseTaskExecutionInformation = original.JobPreparationAndReleaseTaskExecutionInformation
	 JobPreparationTask = original.JobPreparationTask
	 JobPreparationTaskExecutionInformation = original.JobPreparationTaskExecutionInformation
	 JobReleaseTask = original.JobReleaseTask
	 JobReleaseTaskExecutionInformation = original.JobReleaseTaskExecutionInformation
	 JobScheduleAddParameter = original.JobScheduleAddParameter
	 JobScheduleExecutionInformation = original.JobScheduleExecutionInformation
	 JobSchedulePatchParameter = original.JobSchedulePatchParameter
	 JobScheduleStatistics = original.JobScheduleStatistics
	 JobScheduleUpdateParameter = original.JobScheduleUpdateParameter
	 JobSchedulingError = original.JobSchedulingError
	 JobSpecification = original.JobSpecification
	 JobStatistics = original.JobStatistics
	 JobTerminateParameter = original.JobTerminateParameter
	 JobUpdateParameter = original.JobUpdateParameter
	 LinuxUserConfiguration = original.LinuxUserConfiguration
	 MetadataItem = original.MetadataItem
	 MultiInstanceSettings = original.MultiInstanceSettings
	 NameValuePair = original.NameValuePair
	 NetworkConfiguration = original.NetworkConfiguration
	 NetworkSecurityGroupRule = original.NetworkSecurityGroupRule
	 NodeAgentSku = original.NodeAgentSku
	 NodeDisableSchedulingParameter = original.NodeDisableSchedulingParameter
	 NodeFile = original.NodeFile
	 NodeFileListResult = original.NodeFileListResult
	 NodeRebootParameter = original.NodeRebootParameter
	 NodeReimageParameter = original.NodeReimageParameter
	 NodeRemoveParameter = original.NodeRemoveParameter
	 NodeUpdateUserParameter = original.NodeUpdateUserParameter
	 OSDisk = original.OSDisk
	 OutputFile = original.OutputFile
	 OutputFileBlobContainerDestination = original.OutputFileBlobContainerDestination
	 OutputFileDestination = original.OutputFileDestination
	 OutputFileUploadOptions = original.OutputFileUploadOptions
	 PoolAddParameter = original.PoolAddParameter
	 PoolEnableAutoScaleParameter = original.PoolEnableAutoScaleParameter
	 PoolEndpointConfiguration = original.PoolEndpointConfiguration
	 PoolEvaluateAutoScaleParameter = original.PoolEvaluateAutoScaleParameter
	 PoolInformation = original.PoolInformation
	 PoolListUsageMetricsResult = original.PoolListUsageMetricsResult
	 PoolPatchParameter = original.PoolPatchParameter
	 PoolResizeParameter = original.PoolResizeParameter
	 PoolSpecification = original.PoolSpecification
	 PoolStatistics = original.PoolStatistics
	 PoolUpdatePropertiesParameter = original.PoolUpdatePropertiesParameter
	 PoolUpgradeOSParameter = original.PoolUpgradeOSParameter
	 PoolUsageMetrics = original.PoolUsageMetrics
	 ReadCloser = original.ReadCloser
	 RecentJob = original.RecentJob
	 ResizeError = original.ResizeError
	 ResourceFile = original.ResourceFile
	 ResourceStatistics = original.ResourceStatistics
	 Schedule = original.Schedule
	 StartTask = original.StartTask
	 StartTaskInformation = original.StartTaskInformation
	 SubtaskInformation = original.SubtaskInformation
	 TaskAddCollectionParameter = original.TaskAddCollectionParameter
	 TaskAddCollectionResult = original.TaskAddCollectionResult
	 TaskAddParameter = original.TaskAddParameter
	 TaskAddResult = original.TaskAddResult
	 TaskConstraints = original.TaskConstraints
	 TaskCounts = original.TaskCounts
	 TaskDependencies = original.TaskDependencies
	 TaskExecutionInformation = original.TaskExecutionInformation
	 TaskFailureInformation = original.TaskFailureInformation
	 TaskIDRange = original.TaskIDRange
	 TaskInformation = original.TaskInformation
	 TaskSchedulingPolicy = original.TaskSchedulingPolicy
	 TaskStatistics = original.TaskStatistics
	 TaskUpdateParameter = original.TaskUpdateParameter
	 UsageStatistics = original.UsageStatistics
	 UserAccount = original.UserAccount
	 UserIdentity = original.UserIdentity
	 VirtualMachineConfiguration = original.VirtualMachineConfiguration
	 WindowsConfiguration = original.WindowsConfiguration
	 PoolClient = original.PoolClient
	 ManagementClient = original.ManagementClient
	 FileClient = original.FileClient
	 JobClient = original.JobClient
	 JobScheduleClient = original.JobScheduleClient
	 TaskClient = original.TaskClient
)

const (
	 Job = original.Job
	 Resizing = original.Resizing
	 Steady = original.Steady
	 Stopping = original.Stopping
	 Pool = original.Pool
	 Task = original.Task
	 None = original.None
	 ReadOnly = original.ReadOnly
	 ReadWrite = original.ReadWrite
	 Cer = original.Cer
	 Pfx = original.Pfx
	 Active = original.Active
	 DeleteFailed = original.DeleteFailed
	 Deleting = original.Deleting
	 CurrentUser = original.CurrentUser
	 LocalMachine = original.LocalMachine
	 CertificateVisibilityRemoteUser = original.CertificateVisibilityRemoteUser
	 CertificateVisibilityStartTask = original.CertificateVisibilityStartTask
	 CertificateVisibilityTask = original.CertificateVisibilityTask
	 Requeue = original.Requeue
	 RetainedData = original.RetainedData
	 TaskCompletion = original.TaskCompletion
	 Terminate = original.Terminate
	 Pack = original.Pack
	 Spread = original.Spread
	 ComputeNodeRebootOptionRequeue = original.ComputeNodeRebootOptionRequeue
	 ComputeNodeRebootOptionRetainedData = original.ComputeNodeRebootOptionRetainedData
	 ComputeNodeRebootOptionTaskCompletion = original.ComputeNodeRebootOptionTaskCompletion
	 ComputeNodeRebootOptionTerminate = original.ComputeNodeRebootOptionTerminate
	 ComputeNodeReimageOptionRequeue = original.ComputeNodeReimageOptionRequeue
	 ComputeNodeReimageOptionRetainedData = original.ComputeNodeReimageOptionRetainedData
	 ComputeNodeReimageOptionTaskCompletion = original.ComputeNodeReimageOptionTaskCompletion
	 ComputeNodeReimageOptionTerminate = original.ComputeNodeReimageOptionTerminate
	 Creating = original.Creating
	 Idle = original.Idle
	 LeavingPool = original.LeavingPool
	 Offline = original.Offline
	 Preempted = original.Preempted
	 Rebooting = original.Rebooting
	 Reimaging = original.Reimaging
	 Running = original.Running
	 Starting = original.Starting
	 StartTaskFailed = original.StartTaskFailed
	 Unknown = original.Unknown
	 Unusable = original.Unusable
	 WaitingForStartTask = original.WaitingForStartTask
	 Block = original.Block
	 Satisfy = original.Satisfy
	 DisableComputeNodeSchedulingOptionRequeue = original.DisableComputeNodeSchedulingOptionRequeue
	 DisableComputeNodeSchedulingOptionTaskCompletion = original.DisableComputeNodeSchedulingOptionTaskCompletion
	 DisableComputeNodeSchedulingOptionTerminate = original.DisableComputeNodeSchedulingOptionTerminate
	 DisableJobOptionRequeue = original.DisableJobOptionRequeue
	 DisableJobOptionTerminate = original.DisableJobOptionTerminate
	 DisableJobOptionWait = original.DisableJobOptionWait
	 Admin = original.Admin
	 NonAdmin = original.NonAdmin
	 ServerError = original.ServerError
	 UserError = original.UserError
	 TCP = original.TCP
	 UDP = original.UDP
	 JobActionDisable = original.JobActionDisable
	 JobActionNone = original.JobActionNone
	 JobActionTerminate = original.JobActionTerminate
	 JobPreparationTaskStateCompleted = original.JobPreparationTaskStateCompleted
	 JobPreparationTaskStateRunning = original.JobPreparationTaskStateRunning
	 JobReleaseTaskStateCompleted = original.JobReleaseTaskStateCompleted
	 JobReleaseTaskStateRunning = original.JobReleaseTaskStateRunning
	 JobScheduleStateActive = original.JobScheduleStateActive
	 JobScheduleStateCompleted = original.JobScheduleStateCompleted
	 JobScheduleStateDeleting = original.JobScheduleStateDeleting
	 JobScheduleStateDisabled = original.JobScheduleStateDisabled
	 JobScheduleStateTerminating = original.JobScheduleStateTerminating
	 JobStateActive = original.JobStateActive
	 JobStateCompleted = original.JobStateCompleted
	 JobStateDeleting = original.JobStateDeleting
	 JobStateDisabled = original.JobStateDisabled
	 JobStateDisabling = original.JobStateDisabling
	 JobStateEnabling = original.JobStateEnabling
	 JobStateTerminating = original.JobStateTerminating
	 Allow = original.Allow
	 Deny = original.Deny
	 NoAction = original.NoAction
	 TerminateJob = original.TerminateJob
	 OnTaskFailureNoAction = original.OnTaskFailureNoAction
	 OnTaskFailurePerformExitOptionsJobAction = original.OnTaskFailurePerformExitOptionsJobAction
	 Linux = original.Linux
	 Windows = original.Windows
	 OutputFileUploadConditionTaskCompletion = original.OutputFileUploadConditionTaskCompletion
	 OutputFileUploadConditionTaskFailure = original.OutputFileUploadConditionTaskFailure
	 OutputFileUploadConditionTaskSuccess = original.OutputFileUploadConditionTaskSuccess
	 PoolLifetimeOptionJob = original.PoolLifetimeOptionJob
	 PoolLifetimeOptionJobSchedule = original.PoolLifetimeOptionJobSchedule
	 PoolStateActive = original.PoolStateActive
	 PoolStateDeleting = original.PoolStateDeleting
	 PoolStateUpgrading = original.PoolStateUpgrading
	 Disabled = original.Disabled
	 Enabled = original.Enabled
	 StartTaskStateCompleted = original.StartTaskStateCompleted
	 StartTaskStateRunning = original.StartTaskStateRunning
	 SubtaskStateCompleted = original.SubtaskStateCompleted
	 SubtaskStatePreparing = original.SubtaskStatePreparing
	 SubtaskStateRunning = original.SubtaskStateRunning
	 TaskAddStatusClientError = original.TaskAddStatusClientError
	 TaskAddStatusServerError = original.TaskAddStatusServerError
	 TaskAddStatusSuccess = original.TaskAddStatusSuccess
	 Unvalidated = original.Unvalidated
	 Validated = original.Validated
	 Failure = original.Failure
	 Success = original.Success
	 TaskStateActive = original.TaskStateActive
	 TaskStateCompleted = original.TaskStateCompleted
	 TaskStatePreparing = original.TaskStatePreparing
	 TaskStateRunning = original.TaskStateRunning
	 DefaultBaseURI = original.DefaultBaseURI
)