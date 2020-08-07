package main
import (
	"github.com/odwrtw/transmission"
    "errors"
)

type tranmissionplugin struct {
    address string
}
func (tp *tranmissionplugin) Init (options map[string]string){
    tp.address = options["address"]
}
func (tp *tranmissionplugin) getClient() (*transmission.Client, error){
    transmission_config := transmission.Config{
        Address: tp.address,
    }
    client, err := transmission.New(transmission_config)
    _, err = client.GetTorrents()
    if err != nil {
        err = errors.New("failed to connect")
    }
    return client, err
}
func (tp *tranmissionplugin) OnActive() (string, error) {
    /*
    _, err := tp.getClient()
    if err != nil {
        return "", err
    }
    */
    return "success", nil
}
func (tp *tranmissionplugin) OnInactive() (string, error) {
    tc, err := tp.getClient()
    if err != nil {
        return "", err
    }
    torrents, err := tc.GetTorrents()
    if err != nil {
        return "", err
    }
    err = tc.RemoveTorrents(torrents, true)
    if err != nil {
        return "", err
    } else {
        return "success", nil
    }
}
var VpnTwitcherPlugin tranmissionplugin
