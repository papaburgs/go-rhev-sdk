package rhvSDK

// VMs from section 5.238
type VMs struct {
	tonofstuff         string
	path               string
	rawResult          []byte
	paramAllContent    string
	paramCaseSensitive string
	paramFilter        string
	paramMax           string
	paramSearch        string
}

type VMsParams struct {
	AllContent    string
	CaseSensitive string
	Filter        string
	Max           string
	Search        string
}

// Defaults returns sane defaults for params.
// Currently not all content, not case sensitive
// no filter or search and a max of 10
func Defaults() *VMsParams {
	v = VMsParams{
		AllContent:    "false",
		CaseSensitive: "false",
		Filter:        "",
		Max:           "10",
		Search:        "",
	}
	return &v

}

func (v VMs) String() string {
	return string(v.rawResult)
}

func (v *VMs) List(p *VmsParams) {

}

/*
parameters from section 6.266

 bios


Bios


Reference to virtual machine’s BIOS configuration.

comment


String


Free text containing comments about this object.

console


Console


Console configured for this virtual machine.

cpu


Cpu


The configuration of the virtual machine CPU.

cpu_shares


Integer


creation_time


Date


The virtual machine creation date.

custom_compatibility_version


Version


Virtual machine custom compatibility version.

custom_cpu_model


String


custom_emulated_machine


String


custom_properties


CustomProperty[]


Properties sent to VDSM to configure various hooks.

delete_protected


Boolean


If true, the virtual machine cannot be deleted.

description


String


A human-readable description in plain text.

display


Display


The virtual machine display configuration.

domain


Domain


Domain configured for this virtual machine.

fqdn


String


Fully qualified domain name of the virtual machine.

guest_operating_system


GuestOperatingSystem


What operating system is installed on the virtual machine.

guest_time_zone


TimeZone


What time zone is used by the virtual machine (as returned by guest agent).

high_availability


HighAvailability


The virtual machine high availability configuration.

id


String


A unique identifier.

initialization


Initialization


Reference to virtual machine’s initialization configuration.

io


Io


For performance tuning of IO threading.

large_icon


Icon


Virtual machine’s large icon.

memory


Integer


The virtual machine’s memory, in bytes.

memory_policy


MemoryPolicy


Reference to virtual machine’s memory management configuration.

migration


MigrationOptions


Reference to configuration of migration of running virtual machine to another host.

migration_downtime


Integer


Maximum time the virtual machine can be non responsive during its live migration to another host in ms.

name


String


A human-readable name in plain text.

next_run_configuration_exists


Boolean


Virtual machine configuration has been changed and requires restart of the virtual machine.

numa_tune_mode


NumaTuneMode


How the NUMA topology is applied.

origin


String


The origin of this virtual machine.

os


OperatingSystem


Operating system type installed on the virtual machine.

payloads


Payload[]


Optional payloads of the virtual machine, used for ISOs to configure it.

placement_policy


VmPlacementPolicy


The configuration of the virtual machine’s placement policy.

rng_device


RngDevice


Random Number Generator device configuration for this virtual machine.

run_once


Boolean


If true, the virtual machine has been started using the run once command, meaning it’s configuration might differ from the stored one for the purpose of this single run.

serial_number


SerialNumber


Virtual machine’s serial number in a cluster.

small_icon


Icon


Virtual machine’s small icon.

soundcard_enabled


Boolean


If true, the sound card is added to the virtual machine.

sso


Sso


Reference to the Single Sign On configuration this virtual machine is configured for.

start_paused


Boolean


If true, the virtual machine will be initially in 'paused' state after start.

start_time


Date


The date in which the virtual machine was started.

stateless


Boolean


If true, the virtual machine is stateless - it’s state (disks) are rolled-back after shutdown.

status


VmStatus


The current status of the virtual machine.

status_detail


String


Human readable detail of current status.

stop_reason


String


The reason the virtual machine was stopped.

stop_time


Date


The date in which the virtual machine was stopped.

time_zone


TimeZone


The virtual machine’s time zone set by oVirt.

tunnel_migration


Boolean


If true, the network data transfer will be encrypted during virtual machine live migration.

type


VmType


Determines whether the virtual machine is optimized for desktop or server.

usb


Usb


Configuration of USB devices for this virtual machine (count, type).

use_latest_template_version


Boolean


If true, the virtual machine is reconfigured to the latest version of it’s template when it is started.

virtio_scsi


VirtioScsi


Reference to VirtIO SCSI configuration. */
