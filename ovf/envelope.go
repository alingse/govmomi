/*
Copyright (c) 2015-2024 VMware, Inc. All Rights Reserved.

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
	"fmt"
)

// Envelope is defined according to
// https://www.dmtf.org/sites/default/files/standards/documents/DSP0243_2.1.1.pdf.
//
// Section 9 describes the parent/child relationships.
//
// A VirtualSystem may have zero or more VirtualHardware sections.
type Envelope struct {
	References []File `xml:"References>File"`

	// Package level meta-data
	Disk             *DiskSection             `xml:"DiskSection,omitempty"`
	Network          *NetworkSection          `xml:"NetworkSection,omitempty"`
	DeploymentOption *DeploymentOptionSection `xml:"DeploymentOptionSection,omitempty"`

	// Content: A VirtualSystem or a VirtualSystemCollection
	VirtualSystem           *VirtualSystem           `xml:"VirtualSystem,omitempty"`
	VirtualSystemCollection *VirtualSystemCollection `xml:"VirtualSystemCollection,omitempty"`
}

type VirtualSystem struct {
	Content

	Annotation      *AnnotationSection       `xml:"AnnotationSection,omitempty"`
	Product         []ProductSection         `xml:"ProductSection,omitempty"`
	Eula            []EulaSection            `xml:"EulaSection,omitempty"`
	OperatingSystem *OperatingSystemSection  `xml:"OperatingSystemSection,omitempty"`
	VirtualHardware []VirtualHardwareSection `xml:"VirtualHardwareSection,omitempty"`
}

type VirtualSystemCollection struct {
	Content

	// Collection level meta-data
	ResourceAllocation *ResourceAllocationSection `xml:"ResourceAllocationSection,omitempty"`
	Annotation         *AnnotationSection         `xml:"AnnotationSection,omitempty"`
	Product            []ProductSection           `xml:"ProductSection,omitempty"`
	Eula               []EulaSection              `xml:"EulaSection,omitempty"`

	// Content: One or more VirtualSystems
	VirtualSystem []VirtualSystem `xml:"VirtualSystem,omitempty"`
}

type File struct {
	ID          string  `xml:"id,attr"`
	Href        string  `xml:"href,attr"`
	Size        uint    `xml:"size,attr"`
	Compression *string `xml:"compression,attr"`
	ChunkSize   *int    `xml:"chunkSize,attr"`
}

type Content struct {
	ID   string  `xml:"id,attr"`
	Info string  `xml:"Info"`
	Name *string `xml:"Name"`
}

type Section struct {
	Required *bool  `xml:"required,attr"`
	Info     string `xml:"Info"`
	Category string `xml:"Category"`
}

type AnnotationSection struct {
	Section

	Annotation string `xml:"Annotation"`
}

type ProductSection struct {
	Section

	Class    *string `xml:"class,attr"`
	Instance *string `xml:"instance,attr"`

	Product     string     `xml:"Product"`
	Vendor      string     `xml:"Vendor"`
	Version     string     `xml:"Version"`
	FullVersion string     `xml:"FullVersion"`
	ProductURL  string     `xml:"ProductUrl"`
	VendorURL   string     `xml:"VendorUrl"`
	AppURL      string     `xml:"AppUrl"`
	Property    []Property `xml:"Property"`
}

func (p ProductSection) Key(prop Property) string {
	// From OVF spec, section 9.5.1:
	// key-value-env = [class-value "."] key-value-prod ["." instance-value]

	k := prop.Key
	if p.Class != nil {
		k = fmt.Sprintf("%s.%s", *p.Class, k)
	}
	if p.Instance != nil {
		k = fmt.Sprintf("%s.%s", k, *p.Instance)
	}
	return k
}

type Property struct {
	Key              string  `xml:"key,attr"`
	Type             string  `xml:"type,attr"`
	Qualifiers       *string `xml:"qualifiers,attr"`
	UserConfigurable *bool   `xml:"userConfigurable,attr"`
	Default          *string `xml:"value,attr"`
	Password         *bool   `xml:"password,attr"`
	Configuration    *string `xml:"configuration,attr"`

	Label       *string `xml:"Label"`
	Description *string `xml:"Description"`

	Values []PropertyConfigurationValue `xml:"Value"`
}

type PropertyConfigurationValue struct {
	Value         string  `xml:"value,attr"`
	Configuration *string `xml:"configuration,attr"`
}

type NetworkSection struct {
	Section

	Networks []Network `xml:"Network"`
}

type Network struct {
	Name string `xml:"name,attr"`

	Description string `xml:"Description"`
}

type DiskSection struct {
	Section

	Disks []VirtualDiskDesc `xml:"Disk"`
}

type VirtualDiskDesc struct {
	DiskID                  string  `xml:"diskId,attr"`
	FileRef                 *string `xml:"fileRef,attr"`
	Capacity                string  `xml:"capacity,attr"`
	CapacityAllocationUnits *string `xml:"capacityAllocationUnits,attr"`
	Format                  *string `xml:"format,attr"`
	PopulatedSize           *int    `xml:"populatedSize,attr"`
	ParentRef               *string `xml:"parentRef,attr"`
}

type OperatingSystemSection struct {
	Section

	ID      int16   `xml:"id,attr"`
	Version *string `xml:"version,attr"`
	OSType  *string `xml:"osType,attr"`

	Description *string `xml:"Description"`
}

type EulaSection struct {
	Section

	License string `xml:"License"`
}

type Config struct {
	Required *bool  `xml:"required,attr"`
	Key      string `xml:"key,attr"`
	Value    string `xml:"value,attr"`
}

type VirtualHardwareSection struct {
	Section

	ID        *string `xml:"id,attr"`
	Transport *string `xml:"transport,attr"`

	System      *VirtualSystemSettingData       `xml:"System"`
	Item        []ResourceAllocationSettingData `xml:"Item"`
	StorageItem []StorageAllocationSettingData  `xml:"StorageItem"`
	Config      []Config                        `xml:"Config"`
	ExtraConfig []Config                        `xml:"ExtraConfig"`
}

type VirtualSystemSettingData struct {
	CIMVirtualSystemSettingData
}

type ResourceAllocationSettingData struct {
	CIMResourceAllocationSettingData

	Required       *bool           `xml:"required,attr"`
	Configuration  *string         `xml:"configuration,attr"`
	Bound          *string         `xml:"bound,attr"`
	Config         []Config        `xml:"Config"`
	CoresPerSocket *CoresPerSocket `xml:"CoresPerSocket"`
}

type StorageAllocationSettingData struct {
	CIMStorageAllocationSettingData

	Required      *bool   `xml:"required,attr"`
	Configuration *string `xml:"configuration,attr"`
	Bound         *string `xml:"bound,attr"`
}

type ResourceAllocationSection struct {
	Section

	Item []ResourceAllocationSettingData `xml:"Item"`
}

type DeploymentOptionSection struct {
	Section

	Configuration []DeploymentOptionConfiguration `xml:"Configuration"`
}

type DeploymentOptionConfiguration struct {
	ID      string `xml:"id,attr"`
	Default *bool  `xml:"default,attr"`

	Label       string `xml:"Label"`
	Description string `xml:"Description"`
}

type CoresPerSocket struct {
	Required *bool `xml:"required,attr"`
	Value    int32 `xml:",chardata"`
}
