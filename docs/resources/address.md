# b1ddi_address (Resource)

## Example Usage

```terraform
terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "github.com/infobloxopen/b1ddi"
    }
  }
}

resource "b1ddi_ip_space" "example_tf_space" {
  name = "example_tf_space"
  comment = "Example IP space created by the terraform provider"
}

resource "b1ddi_subnet" "example_tf_subnet" {
  name = "example_tf_subnet"
  space = b1ddi_ip_space.example_tf_space.id
  address = "192.168.1.0"
  cidr = 24
  comment = "Example Subnet created by the terraform provider"
}

resource "b1ddi_address" "example_tf_address" {
  address = "192.168.1.45"
  comment = "Example Address created by the terraform provider"
  space = b1ddi_ip_space.example_tf_space.id
  depends_on = [b1ddi_subnet.example_tf_subnet]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **address** (String) The address in form "a.b.c.d".
- **space** (String) The resource identifier.

### Optional

- **comment** (String) The description for the address object. May contain 0 to 1024 characters. Can include UTF-8.
- **host** (String) The resource identifier.
- **hwaddr** (String) The hardware address associated with this IP address.
- **id** (String) The ID of this resource.
- **interface** (String) The name of the network interface card (NIC) associated with the address, if any.
- **names** (Block List) The list of all names associated with this address. (see [below for nested schema](#nestedblock--names))
- **parent** (String) The resource identifier.
- **range** (String) The resource identifier.
- **tags** (Map of String) The tags for this address in JSON format.

### Read-Only

- **created_at** (String) Time when the object has been created.
- **dhcp_info** (List of Object) The DHCP information associated with this object. (see [below for nested schema](#nestedatt--dhcp_info))
- **protocol** (String) The type of protocol (_ipv4_ or _ipv6_).
- **state** (String) The state of the address (_used_ or _free_).
- **updated_at** (String) Time when the object has been updated. Equals to _created_at_ if not updated after creation.
- **usage** (List of String) The usage is a combination of indicators, each tracking a specific associated use. Listed below are usage indicators with their meaning:
 
  | usage indicator | description |
  | --- | --- |
  | _IPAM_ | Address was created by the IPAM component. |
  | _IPAM_, _RESERVED_ | Address was created by the API call _ipam/address_ or _ipam/host_. | 
  | _IPAM_, _NETWORK_ | Address was automatically created by the IPAM component and is the network address of the parent subnet. |
  | _IPAM_, _BROADCAST_ | Address was automatically created by the IPAM component and is the broadcast address of the parent subnet. |
  | _DHCP_ | Address was created by the DHCP component. |
  | _DHCP_, _FIXEDADDRESS_ | Address was created by the API call _dhcp/fixed_address_. |
  | _DHCP_, _LEASED_ | An active lease for that address was issued by a DHCP server. |
  | _DNS_ | Address is used by one or more DNS records. |

<a id="nestedblock--names"></a>
<details>
<summary>Nested Schema for `names`</summary> 

Required:

- **name** (String) The name expressed as a single label or FQDN.
- **type** (String) The origin of the name.

</details>

<a id="nestedatt--dhcp_info"></a>
<details>
<summary>Nested Schema for `dhcp_info`</summary>

Read-Only:

- **client_hostname** (String) The DHCP host name associated with this client.
- **client_hwaddr** (String) The hardware address associated with this client.
- **client_id** (String) The ID associated with this client.
- **end** (String) The timestamp at which the _state_, when set to _leased_, will be changed to _free_.
- **fingerprint** (String) The DHCP fingerprint for the associated lease.
- **remain** (Number) The remaining time, in seconds, until which the _state_, when set to _leased_, will remain in that state.
- **start** (String) The timestamp at which _state_ was first set to _leased_.
- **state** (String) Indicates the status of this IP address from a DHCP protocol standpoint as:
  * _none_: Address is not under DHCP control.
  * _free_: Address is under DHCP control but has no lease currently assigned.
  * _leased_: Address is under DHCP control and has a lease currently assigned. The lease details are contained in the matching _dhcp/lease_ resource.
- **state_ts** (String) The timestamp at which the _state_ was last reported.

</details>
