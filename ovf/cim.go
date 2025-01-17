/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ovf

import (
	"github.com/vmware/govmomi/vim25/types"
)

/*
Source:
  - https://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.24.0/CIM_ResourceAllocationSettingData.xsd
  - https://schemas.dmtf.org/wbem/cim-html/2/CIM_ResourceAllocationSettingData.html
*/
type CIMResourceType uint16

// Please note, the iota pattern is not used to ensure these constants remain
// affixed to an explicit value.
const (
	Other              CIMResourceType = 1
	ComputerSystem     CIMResourceType = 2
	Processor          CIMResourceType = 3
	Memory             CIMResourceType = 4
	IdeController      CIMResourceType = 5
	ParallelScsiHba    CIMResourceType = 6
	FcHba              CIMResourceType = 7
	IScsiHba           CIMResourceType = 8
	IbHba              CIMResourceType = 9
	EthernetAdapter    CIMResourceType = 10
	OtherNetwork       CIMResourceType = 11
	IoSlot             CIMResourceType = 12
	IoDevice           CIMResourceType = 13
	FloppyDrive        CIMResourceType = 14
	CdDrive            CIMResourceType = 15
	DvdDrive           CIMResourceType = 16
	DiskDrive          CIMResourceType = 17
	TapeDrive          CIMResourceType = 18
	StorageExtent      CIMResourceType = 19
	OtherStorage       CIMResourceType = 20
	SerialPort         CIMResourceType = 21
	ParallelPort       CIMResourceType = 22
	UsbController      CIMResourceType = 23
	Graphics           CIMResourceType = 24
	Ieee1394           CIMResourceType = 25
	PartitionableUnit  CIMResourceType = 26
	BasePartitionable  CIMResourceType = 27
	PowerSupply        CIMResourceType = 28
	CoolingDevice      CIMResourceType = 29
	EthernetSwitchPort CIMResourceType = 30
	LogicalDisk        CIMResourceType = 31
	StorageVolume      CIMResourceType = 32
	EthernetConnection CIMResourceType = 33
)

/*
Source: http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.24.0/CIM_VirtualSystemSettingData.xsd
*/

type CIMVirtualSystemSettingData struct {
	ElementName string `xml:"ElementName"`
	InstanceID  string `xml:"InstanceID"`

	AutomaticRecoveryAction              *uint8   `xml:"AutomaticRecoveryAction"`
	AutomaticShutdownAction              *uint8   `xml:"AutomaticShutdownAction"`
	AutomaticStartupAction               *uint8   `xml:"AutomaticStartupAction"`
	AutomaticStartupActionDelay          *string  `xml:"AutomaticStartupActionDelay>Interval"`
	AutomaticStartupActionSequenceNumber *uint16  `xml:"AutomaticStartupActionSequenceNumber"`
	Caption                              *string  `xml:"Caption"`
	ConfigurationDataRoot                *string  `xml:"ConfigurationDataRoot"`
	ConfigurationFile                    *string  `xml:"ConfigurationFile"`
	ConfigurationID                      *string  `xml:"ConfigurationID"`
	CreationTime                         *string  `xml:"CreationTime"`
	Description                          *string  `xml:"Description"`
	LogDataRoot                          *string  `xml:"LogDataRoot"`
	Notes                                []string `xml:"Notes"`
	RecoveryFile                         *string  `xml:"RecoveryFile"`
	SnapshotDataRoot                     *string  `xml:"SnapshotDataRoot"`
	SuspendDataRoot                      *string  `xml:"SuspendDataRoot"`
	SwapFileDataRoot                     *string  `xml:"SwapFileDataRoot"`
	VirtualSystemIdentifier              *string  `xml:"VirtualSystemIdentifier"`
	VirtualSystemType                    *string  `xml:"VirtualSystemType"`
}

/*
Source: http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.24.0/CIM_ResourceAllocationSettingData.xsd
*/

type CIMResourceAllocationSettingData struct {
	ElementName string `xml:"ElementName"`
	InstanceID  string `xml:"InstanceID"`

	ResourceType      *CIMResourceType `xml:"ResourceType"`
	OtherResourceType *string          `xml:"OtherResourceType"`
	ResourceSubType   *string          `xml:"ResourceSubType"`

	AddressOnParent       *string  `xml:"AddressOnParent"`
	Address               *string  `xml:"Address"`
	AllocationUnits       *string  `xml:"AllocationUnits"`
	AutomaticAllocation   *bool    `xml:"AutomaticAllocation"`
	AutomaticDeallocation *bool    `xml:"AutomaticDeallocation"`
	Caption               *string  `xml:"Caption"`
	Connection            []string `xml:"Connection"`
	ConsumerVisibility    *uint16  `xml:"ConsumerVisibility"`
	Description           *string  `xml:"Description"`
	HostResource          []string `xml:"HostResource"`
	Limit                 *uint64  `xml:"Limit"`
	MappingBehavior       *uint    `xml:"MappingBehavior"`
	Parent                *string  `xml:"Parent"`
	PoolID                *string  `xml:"PoolID"`
	Reservation           *uint64  `xml:"Reservation"`
	VirtualQuantity       *uint    `xml:"VirtualQuantity"`
	VirtualQuantityUnits  *string  `xml:"VirtualQuantityUnits"`
	Weight                *uint    `xml:"Weight"`
}

/*
Source: http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.24.0/CIM_StorageAllocationSettingData.xsd
*/
type CIMStorageAllocationSettingData struct {
	ElementName string `xml:"ElementName"`
	InstanceID  string `xml:"InstanceID"`

	ResourceType      *CIMResourceType `xml:"ResourceType"`
	OtherResourceType *string          `xml:"OtherResourceType"`
	ResourceSubType   *string          `xml:"ResourceSubType"`

	Access                       *uint16         `xml:"Access"`
	Address                      *string         `xml:"Address"`
	AddressOnParent              *string         `xml:"AddressOnParent"`
	AllocationUnits              *string         `xml:"AllocationUnits"`
	AutomaticAllocation          *bool           `xml:"AutomaticAllocation"`
	AutomaticDeallocation        *bool           `xml:"AutomaticDeallocation"`
	Caption                      *string         `xml:"Caption"`
	ChangeableType               *uint16         `xml:"ChangeableType"`
	ComponentSetting             []types.AnyType `xml:"ComponentSetting"`
	ConfigurationName            *string         `xml:"ConfigurationName"`
	Connection                   []string        `xml:"Connection"`
	ConsumerVisibility           *uint16         `xml:"ConsumerVisibility"`
	Description                  *string         `xml:"Description"`
	Generation                   *uint64         `xml:"Generation"`
	HostExtentName               *string         `xml:"HostExtentName"`
	HostExtentNameFormat         *uint16         `xml:"HostExtentNameFormat"`
	HostExtentNameNamespace      *uint16         `xml:"HostExtentNameNamespace"`
	HostExtentStartingAddress    *uint64         `xml:"HostExtentStartingAddress"`
	HostResource                 []string        `xml:"HostResource"`
	HostResourceBlockSize        *uint64         `xml:"HostResourceBlockSize"`
	Limit                        *uint64         `xml:"Limit"`
	MappingBehavior              *uint           `xml:"MappingBehavior"`
	OtherHostExtentNameFormat    *string         `xml:"OtherHostExtentNameFormat"`
	OtherHostExtentNameNamespace *string         `xml:"OtherHostExtentNameNamespace"`
	Parent                       *string         `xml:"Parent"`
	PoolID                       *string         `xml:"PoolID"`
	Reservation                  *uint64         `xml:"Reservation"`
	SoID                         *string         `xml:"SoID"`
	SoOrgID                      *string         `xml:"SoOrgID"`
	VirtualQuantity              *uint           `xml:"VirtualQuantity"`
	VirtualQuantityUnits         *string         `xml:"VirtualQuantityUnits"`
	VirtualResourceBlockSize     *uint64         `xml:"VirtualResourceBlockSize"`
	Weight                       *uint           `xml:"Weight"`
}
