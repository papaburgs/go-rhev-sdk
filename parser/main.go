package main

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/html"
)

func main() {
	z := html.NewTokenizer(strings.NewReader(cpu_table))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	i := 0
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return
		}
		if 
		fmt.Println("--------------")
		fmt.Println(tt.String())
		t := z.Token()
		spew.Dump(t)
		i++
	}
}



var cpu_table = `<section class="section" id="types-cpu"><div class="titlepage"><div><div><h2 class="title">6.33.&nbsp;Cpu <span class="small small">struct</span></h2></div></div></div><div class="table" id="idm140613544648832"><p class="title"><strong>Table&nbsp;6.42.&nbsp;Attributes summary</strong></p><div class="table-contents"><table class="lt-4-cols lt-7-rows"><colgroup><col style="width: 20%; " class="col_1"><!--Empty--><col style="width: 20%; " class="col_2"><!--Empty--><col style="width: 60%; " class="col_3"><!--Empty--></colgroup><thead><tr><th style="text-align: left; vertical-align: top; ">Name</th><th style="text-align: left; vertical-align: top; ">Type</th><th style="text-align: left; vertical-align: top; ">Summary</th></tr></thead><tbody><tr><td style="text-align: left; vertical-align: top; "> <p>
								<code class="literal">architecture</code>
							</p>
							 </td><td style="text-align: left; vertical-align: top; "> <p>
								<a class="link" href="#types-architecture" title="6.11.&nbsp;Architecture enum">Architecture</a>
							</p>
							 </td><td style="text-align: left; vertical-align: top; ">&nbsp;</td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
								<code class="literal">cores</code>
							</p>
							 </td><td style="text-align: left; vertical-align: top; "> <p>
								<a class="link" href="#types-core" title="6.32.&nbsp;Core struct">Core[]</a>
							</p>
							 </td><td style="text-align: left; vertical-align: top; ">&nbsp;</td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
								<code class="literal">cpu_tune</code>
							</p>
							 </td><td style="text-align: left; vertical-align: top; "> <p>
								<a class="link" href="#types-cpu_tune" title="6.37.&nbsp;CpuTune struct">CpuTune</a>
							</p>
							 </td><td style="text-align: left; vertical-align: top; ">&nbsp;</td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
								<code class="literal">level</code>
							</p>
							 </td><td style="text-align: left; vertical-align: top; "> <p>
								<a class="link" href="#types-integer" title="A.3.&nbsp;Integer primitive">Integer</a>
							</p>
							 </td><td style="text-align: left; vertical-align: top; ">&nbsp;</td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
								<code class="literal">mode</code>
							</p>
							 </td><td style="text-align: left; vertical-align: top; "> <p>
								<a class="link" href="#types-cpu_mode" title="6.34.&nbsp;CpuMode enum">CpuMode</a>
							</p>
							 </td><td style="text-align: left; vertical-align: top; ">&nbsp;</td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
								<code class="literal">name</code>
							</p>
							 </td><td style="text-align: left; vertical-align: top; "> <p>
								<a class="link" href="#types-string" title="A.1.&nbsp;String primitive">String</a>
							</p>
							 </td><td style="text-align: left; vertical-align: top; ">&nbsp;</td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
								<code class="literal">speed</code>
							</p>
							 </td><td style="text-align: left; vertical-align: top; "> <p>
								<a class="link" href="#types-decimal" title="A.4.&nbsp;Decimal primitive">Decimal</a>
							</p>
							 </td><td style="text-align: left; vertical-align: top; ">&nbsp;</td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
								<code class="literal">topology</code>
							</p>
							 </td><td style="text-align: left; vertical-align: top; "> <p>
								<a class="link" href="#types-cpu_topology" title="6.36.&nbsp;CpuTopology struct">CpuTopology</a>
							</p>
							 </td><td style="text-align: left; vertical-align: top; ">&nbsp;</td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
								<code class="literal">type</code>
							</p>
							 </td><td style="text-align: left; vertical-align: top; "> <p>
								<a class="link" href="#types-string" title="A.1.&nbsp;String primitive">String</a>
							</p>
							 </td><td style="text-align: left; vertical-align: top; ">&nbsp;</td></tr></tbody></table></div></div></section>`

var table = `<p class="title"><strong>Table&nbsp;6.347.&nbsp;Attributes summary</strong></p><div class="table-contents"><table class="lt-4-cols lt-7-rows"><colgroup><col style="width: 20%; " class="col_1"><!--Empty--><col style="width: 20%; " class="col_2"><!--Empty--><col style="width: 60%; " class="col_3"><!--Empty--></colgroup><thead><tr><th style="text-align: left; vertical-align: top; ">Name</th><th style="text-align: left; vertical-align: top; ">Type</th><th style="text-align: left; vertical-align: top; ">Summary</th></tr></thead><tbody><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">bios</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-bios" title="6.15.&nbsp;Bios struct">Bios</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Reference to virtual machine’s BIOS configuration.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">comment</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-string" title="A.1.&nbsp;String primitive">String</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Free text containing comments about this object.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">console</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-console" title="6.31.&nbsp;Console struct">Console</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Console configured for this virtual machine.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">cpu</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-cpu" title="6.33.&nbsp;Cpu struct">Cpu</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							The configuration of the virtual machine CPU.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">cpu_shares</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-integer" title="A.3.&nbsp;Integer primitive">Integer</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; ">&nbsp;</td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">creation_time</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-date" title="A.5.&nbsp;Date primitive">Date</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							The virtual machine creation date.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">custom_compatibility_version</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-version" title="6.262.&nbsp;Version struct">Version</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Virtual machine custom compatibility version.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">custom_cpu_model</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-string" title="A.1.&nbsp;String primitive">String</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; ">&nbsp;</td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">custom_emulated_machine</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-string" title="A.1.&nbsp;String primitive">String</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; ">&nbsp;</td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">custom_properties</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-custom_property" title="6.40.&nbsp;CustomProperty struct">CustomProperty[]</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Properties sent to VDSM to configure various hooks.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">delete_protected</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-boolean" title="A.2.&nbsp;Boolean primitive">Boolean</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							If <code class="literal">true</code>, the virtual machine cannot be deleted.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">description</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-string" title="A.1.&nbsp;String primitive">String</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							A human-readable description in plain text.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">display</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-display" title="6.53.&nbsp;Display struct">Display</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							The virtual machine display configuration.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">domain</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-domain" title="6.56.&nbsp;Domain struct">Domain</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Domain configured for this virtual machine.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">fqdn</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-string" title="A.1.&nbsp;String primitive">String</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Fully qualified domain name of the virtual machine.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">guest_operating_system</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-guest_operating_system" title="6.96.&nbsp;GuestOperatingSystem struct">GuestOperatingSystem</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							What operating system is installed on the virtual machine.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">guest_time_zone</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-time_zone" title="6.251.&nbsp;TimeZone struct">TimeZone</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							What time zone is used by the virtual machine (as returned by guest agent).
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">high_availability</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-high_availability" title="6.98.&nbsp;HighAvailability struct">HighAvailability</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							The virtual machine high availability configuration.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">id</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-string" title="A.1.&nbsp;String primitive">String</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							A unique identifier.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">initialization</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-initialization" title="6.119.&nbsp;Initialization struct">Initialization</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Reference to virtual machine’s initialization configuration.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">io</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-io" title="6.121.&nbsp;Io struct">Io</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							For performance tuning of IO threading.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">large_icon</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-icon" title="6.113.&nbsp;Icon struct">Icon</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Virtual machine’s large icon.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">memory</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-integer" title="A.3.&nbsp;Integer primitive">Integer</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							The virtual machine’s memory, in bytes.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">memory_policy</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-memory_policy" title="6.139.&nbsp;MemoryPolicy struct">MemoryPolicy</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Reference to virtual machine’s memory management configuration.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">migration</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-migration_options" title="6.145.&nbsp;MigrationOptions struct">MigrationOptions</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Reference to configuration of migration of running virtual machine to another host.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">migration_downtime</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-integer" title="A.3.&nbsp;Integer primitive">Integer</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Maximum time the virtual machine can be non responsive during its live migration to another host in ms.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">name</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-string" title="A.1.&nbsp;String primitive">String</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							A human-readable name in plain text.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">next_run_configuration_exists</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-boolean" title="A.2.&nbsp;Boolean primitive">Boolean</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Virtual machine configuration has been changed and requires restart of the virtual machine.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">numa_tune_mode</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-numa_tune_mode" title="6.163.&nbsp;NumaTuneMode enum">NumaTuneMode</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							How the NUMA topology is applied.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">origin</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-string" title="A.1.&nbsp;String primitive">String</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							The origin of this virtual machine.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">os</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-operating_system" title="6.175.&nbsp;OperatingSystem struct">OperatingSystem</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Operating system type installed on the virtual machine.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">payloads</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-payload" title="6.180.&nbsp;Payload struct">Payload[]</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Optional payloads of the virtual machine, used for ISOs to configure it.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">placement_policy</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-vm_placement_policy" title="6.270.&nbsp;VmPlacementPolicy struct">VmPlacementPolicy</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							The configuration of the virtual machine’s placement policy.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">rng_device</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-rng_device" title="6.207.&nbsp;RngDevice struct">RngDevice</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Random Number Generator device configuration for this virtual machine.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">run_once</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-boolean" title="A.2.&nbsp;Boolean primitive">Boolean</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							If <code class="literal">true</code>, the virtual machine has been started using the <span class="emphasis"><em>run once</em></span> command, meaning it’s configuration might differ from the stored one for the purpose of this single run.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">serial_number</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-serial_number" title="6.216.&nbsp;SerialNumber struct">SerialNumber</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Virtual machine’s serial number in a cluster.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">small_icon</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-icon" title="6.113.&nbsp;Icon struct">Icon</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Virtual machine’s small icon.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">soundcard_enabled</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-boolean" title="A.2.&nbsp;Boolean primitive">Boolean</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							If <code class="literal">true</code>, the sound card is added to the virtual machine.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">sso</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-sso" title="6.230.&nbsp;Sso struct">Sso</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Reference to the Single Sign On configuration this virtual machine is configured for.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">start_paused</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-boolean" title="A.2.&nbsp;Boolean primitive">Boolean</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							If <code class="literal">true</code>, the virtual machine will be initially in 'paused' state after start.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">start_time</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-date" title="A.5.&nbsp;Date primitive">Date</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							The date in which the virtual machine was started.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">stateless</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-boolean" title="A.2.&nbsp;Boolean primitive">Boolean</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							If <code class="literal">true</code>, the virtual machine is stateless - it’s state (disks) are rolled-back after shutdown.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">status</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-vm_status" title="6.273.&nbsp;VmStatus enum">VmStatus</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							The current status of the virtual machine.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">status_detail</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-string" title="A.1.&nbsp;String primitive">String</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Human readable detail of current status.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">stop_reason</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-string" title="A.1.&nbsp;String primitive">String</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							The reason the virtual machine was stopped.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">stop_time</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-date" title="A.5.&nbsp;Date primitive">Date</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							The date in which the virtual machine was stopped.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">time_zone</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-time_zone" title="6.251.&nbsp;TimeZone struct">TimeZone</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							The virtual machine’s time zone set by oVirt.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">tunnel_migration</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-boolean" title="A.2.&nbsp;Boolean primitive">Boolean</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							If <code class="literal">true</code>, the network data transfer will be encrypted during virtual machine live migration.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">type</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-vm_type" title="6.275.&nbsp;VmType enum">VmType</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Determines whether the virtual machine is optimized for desktop or server.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">usb</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-usb" title="6.255.&nbsp;Usb struct">Usb</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Configuration of USB devices for this virtual machine (count, type).
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">use_latest_template_version</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-boolean" title="A.2.&nbsp;Boolean primitive">Boolean</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							If <code class="literal">true</code>, the virtual machine is reconfigured to the latest version of it’s template when it is started.
						</p>
						 </td></tr><tr><td style="text-align: left; vertical-align: top; "> <p>
							<code class="literal">virtio_scsi</code>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							<a class="link" href="#types-virtio_scsi" title="6.263.&nbsp;VirtioScsi struct">VirtioScsi</a>
						</p>
						 </td><td style="text-align: left; vertical-align: top; "> <p>
							Reference to VirtIO SCSI configuration.
						</p>
						 </td></tr></tbody></table></div> `
