package main
import (
	"github.com/odwrtw/transmission"
    "fmt"
)

type tranmissionplugin struct {
    address string
}
func (tp *tranmissionplugin) Init (options map[string]string){
    tp.address = options["address"]
}
func (tp *tranmissionplugin) getClient() (*transmission.Client){
    transmission_config := transmission.Config{
        Address: tp.address,
    }
    client, err = transmission.New(transmission_config)
}
func (tp *tranmissionplugin) OnActive() {
}
func (tp *tranmissionplugin) OnInactive() {
}
var VpnTwitcherPlugin tranmissionplugin