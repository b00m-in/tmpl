package tmpl

import (
)

type PostFormData struct {
        Start string `json:"start,omitemtpy" form:"start"`
        End string `json:"end,omitempty" form:"end"`
        Freq string `json:"freq,omitempty" form:"freq"`
}

type PostFormPubConfig struct {
        Kwp string `json:"kwp,omitemtpy" form:"kwp"`
        Kwpmake string `json:"kwpm,omitempty" form:"kwpm"`
        Kwr string `json:"kwr,omitempty" form:"kwr"`
        Kwrmake string `json:"kwrm,omitempty" form:"kwrm"`
        Notify string `json:"notify,omitempty" form:"notify"`
}

type PostFormSub struct {
        Old string `json:"old_password,omitemtpy" form:"old_password"`
        Pswd string `json:"password,omitempty" form:"password"`
        Conf string `json:"confirm,omitempty" form:"confirm"`
}
