package main
import (
	"testing"
)

func TestConnection(t *testing.T){
    var options = map[string] string {"address": ""}
    VpnTwitcherPlugin.Init(options)
    _, err := VpnTwitcherPlugin.OnActive()
    if err == nil {
        t.Error("the connection should fail")
    }
    _, err = VpnTwitcherPlugin.OnInactive()
    if err == nil {
        t.Error("the connection should fail")
    }
    options["address"] = "http://10.140.10.70:9091/transmission/rpc"
    VpnTwitcherPlugin.Init(options)
    _, err = VpnTwitcherPlugin.OnActive()
    if err != nil {
        t.Error("the connection should pass")
    }
    _, err = VpnTwitcherPlugin.OnInactive()
    if err != nil {
        t.Error("the connection should pass")
    }
}