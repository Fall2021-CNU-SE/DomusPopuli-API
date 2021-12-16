package service

import (
	"strings"

	"github.com/NamSoGong/DomusPopuli-API/domain"
	"github.com/NamSoGong/DomusPopuli-API/repository"
	"github.com/NamSoGong/DomusPopuli-API/domain/api"
)

func UpdatePreferences(sid uint, pref api.Preference_t) error {
    var err error

    //  Handling workspace address
    var coords *domain.Coordinate_t = nil
    if pref.Workspace != "" {
        if coords, err = AddressToCoordinate(pref.Workspace); err != nil {
            return err
        }
    }

    facs := strings.Join(pref.Facilities, "/")
    prefs := PrefWeight(pref.Favors)

    if err := repository.UpdatePreferences(sid, pref.Budget, coords, facs, *prefs); err != nil {
        return err
    }

    return nil
}
