// Auto-generated by avdl-compiler v1.3.17 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/teams.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type TeamRole int

const (
	TeamRole_NONE   TeamRole = 0
	TeamRole_OWNER  TeamRole = 1
	TeamRole_ADMIN  TeamRole = 2
	TeamRole_WRITER TeamRole = 3
	TeamRole_READER TeamRole = 4
)

func (o TeamRole) DeepCopy() TeamRole { return o }

var TeamRoleMap = map[string]TeamRole{
	"NONE":   0,
	"OWNER":  1,
	"ADMIN":  2,
	"WRITER": 3,
	"READER": 4,
}

var TeamRoleRevMap = map[TeamRole]string{
	0: "NONE",
	1: "OWNER",
	2: "ADMIN",
	3: "WRITER",
	4: "READER",
}

func (e TeamRole) String() string {
	if v, ok := TeamRoleRevMap[e]; ok {
		return v
	}
	return ""
}

type TeamApplication int

const (
	TeamApplication_KBFS     TeamApplication = 1
	TeamApplication_CHAT     TeamApplication = 2
	TeamApplication_SALTPACK TeamApplication = 3
)

func (o TeamApplication) DeepCopy() TeamApplication { return o }

var TeamApplicationMap = map[string]TeamApplication{
	"KBFS":     1,
	"CHAT":     2,
	"SALTPACK": 3,
}

var TeamApplicationRevMap = map[TeamApplication]string{
	1: "KBFS",
	2: "CHAT",
	3: "SALTPACK",
}

func (e TeamApplication) String() string {
	if v, ok := TeamApplicationRevMap[e]; ok {
		return v
	}
	return ""
}

type TeamApplicationKey struct {
	Application   TeamApplication `codec:"application" json:"application"`
	KeyGeneration int             `codec:"keyGeneration" json:"keyGeneration"`
	Key           Bytes32         `codec:"key" json:"key"`
}

func (o TeamApplicationKey) DeepCopy() TeamApplicationKey {
	return TeamApplicationKey{
		Application:   o.Application.DeepCopy(),
		KeyGeneration: o.KeyGeneration,
		Key:           o.Key.DeepCopy(),
	}
}

type MaskB64 []byte

func (o MaskB64) DeepCopy() MaskB64 {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte(nil), x...)
	})(o)
}

type ReaderKeyMask struct {
	Application TeamApplication `codec:"application" json:"application"`
	Generation  int             `codec:"generation" json:"generation"`
	Mask        MaskB64         `codec:"mask" json:"mask"`
}

func (o ReaderKeyMask) DeepCopy() ReaderKeyMask {
	return ReaderKeyMask{
		Application: o.Application.DeepCopy(),
		Generation:  o.Generation,
		Mask:        o.Mask.DeepCopy(),
	}
}

type PerTeamKey struct {
	Gen    int   `codec:"gen" json:"gen"`
	Seqno  Seqno `codec:"seqno" json:"seqno"`
	SigKID KID   `codec:"sigKID" json:"sigKID"`
	EncKID KID   `codec:"encKID" json:"encKID"`
}

func (o PerTeamKey) DeepCopy() PerTeamKey {
	return PerTeamKey{
		Gen:    o.Gen,
		Seqno:  o.Seqno.DeepCopy(),
		SigKID: o.SigKID.DeepCopy(),
		EncKID: o.EncKID.DeepCopy(),
	}
}

type TeamMembers struct {
	Owners  []string `codec:"owners" json:"owners"`
	Admins  []string `codec:"admins" json:"admins"`
	Writers []string `codec:"writers" json:"writers"`
	Readers []string `codec:"readers" json:"readers"`
}

func (o TeamMembers) DeepCopy() TeamMembers {
	return TeamMembers{
		Owners: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Owners),
		Admins: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Admins),
		Writers: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Writers),
		Readers: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Readers),
	}
}

type TeamChangeReq struct {
	Owners  []string `codec:"owners" json:"owners"`
	Admins  []string `codec:"admins" json:"admins"`
	Writers []string `codec:"writers" json:"writers"`
	Readers []string `codec:"readers" json:"readers"`
	None    []string `codec:"none" json:"none"`
}

func (o TeamChangeReq) DeepCopy() TeamChangeReq {
	return TeamChangeReq{
		Owners: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Owners),
		Admins: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Admins),
		Writers: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Writers),
		Readers: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Readers),
		None: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.None),
	}
}

type UserVersion struct {
	Username    string `codec:"username" json:"username"`
	EldestSeqno Seqno  `codec:"eldestSeqno" json:"eldestSeqno"`
}

func (o UserVersion) DeepCopy() UserVersion {
	return UserVersion{
		Username:    o.Username,
		EldestSeqno: o.EldestSeqno.DeepCopy(),
	}
}

type TeamSigChainState struct {
	Reader       UserVersion                    `codec:"reader" json:"reader"`
	Id           TeamID                         `codec:"id" json:"id"`
	Name         string                         `codec:"name" json:"name"`
	LastSeqno    Seqno                          `codec:"lastSeqno" json:"lastSeqno"`
	LastLinkID   LinkID                         `codec:"lastLinkID" json:"lastLinkID"`
	ParentID     *TeamID                        `codec:"parentID,omitempty" json:"parentID,omitempty"`
	UserLog      map[UserVersion][]UserLogPoint `codec:"userLog" json:"userLog"`
	PerTeamKeys  map[int]PerTeamKey             `codec:"perTeamKeys" json:"perTeamKeys"`
	StubbedTypes map[int]bool                   `codec:"stubbedTypes" json:"stubbedTypes"`
}

func (o TeamSigChainState) DeepCopy() TeamSigChainState {
	return TeamSigChainState{
		Reader:     o.Reader.DeepCopy(),
		Id:         o.Id.DeepCopy(),
		Name:       o.Name,
		LastSeqno:  o.LastSeqno.DeepCopy(),
		LastLinkID: o.LastLinkID.DeepCopy(),
		ParentID: (func(x *TeamID) *TeamID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ParentID),
		UserLog: (func(x map[UserVersion][]UserLogPoint) map[UserVersion][]UserLogPoint {
			ret := make(map[UserVersion][]UserLogPoint)
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x []UserLogPoint) []UserLogPoint {
					var ret []UserLogPoint
					for _, v := range x {
						vCopy := v.DeepCopy()
						ret = append(ret, vCopy)
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.UserLog),
		PerTeamKeys: (func(x map[int]PerTeamKey) map[int]PerTeamKey {
			ret := make(map[int]PerTeamKey)
			for k, v := range x {
				kCopy := k
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PerTeamKeys),
		StubbedTypes: (func(x map[int]bool) map[int]bool {
			ret := make(map[int]bool)
			for k, v := range x {
				kCopy := k
				vCopy := v
				ret[kCopy] = vCopy
			}
			return ret
		})(o.StubbedTypes),
	}
}

type UserLogPoint struct {
	Role  TeamRole `codec:"role" json:"role"`
	Seqno Seqno    `codec:"seqno" json:"seqno"`
}

func (o UserLogPoint) DeepCopy() UserLogPoint {
	return UserLogPoint{
		Role:  o.Role.DeepCopy(),
		Seqno: o.Seqno.DeepCopy(),
	}
}

type TeamCreateArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Name      string `codec:"name" json:"name"`
}

func (o TeamCreateArg) DeepCopy() TeamCreateArg {
	return TeamCreateArg{
		SessionID: o.SessionID,
		Name:      o.Name,
	}
}

type TeamGetArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Name      string `codec:"name" json:"name"`
}

func (o TeamGetArg) DeepCopy() TeamGetArg {
	return TeamGetArg{
		SessionID: o.SessionID,
		Name:      o.Name,
	}
}

type TeamChangeMembershipArg struct {
	SessionID int           `codec:"sessionID" json:"sessionID"`
	Name      string        `codec:"name" json:"name"`
	Req       TeamChangeReq `codec:"req" json:"req"`
}

func (o TeamChangeMembershipArg) DeepCopy() TeamChangeMembershipArg {
	return TeamChangeMembershipArg{
		SessionID: o.SessionID,
		Name:      o.Name,
		Req:       o.Req.DeepCopy(),
	}
}

type TeamAddMemberArg struct {
	SessionID            int      `codec:"sessionID" json:"sessionID"`
	Name                 string   `codec:"name" json:"name"`
	Username             string   `codec:"username" json:"username"`
	Role                 TeamRole `codec:"role" json:"role"`
	SendChatNotification bool     `codec:"sendChatNotification" json:"sendChatNotification"`
}

func (o TeamAddMemberArg) DeepCopy() TeamAddMemberArg {
	return TeamAddMemberArg{
		SessionID:            o.SessionID,
		Name:                 o.Name,
		Username:             o.Username,
		Role:                 o.Role.DeepCopy(),
		SendChatNotification: o.SendChatNotification,
	}
}

type TeamRemoveMemberArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Name      string `codec:"name" json:"name"`
	Username  string `codec:"username" json:"username"`
}

func (o TeamRemoveMemberArg) DeepCopy() TeamRemoveMemberArg {
	return TeamRemoveMemberArg{
		SessionID: o.SessionID,
		Name:      o.Name,
		Username:  o.Username,
	}
}

type TeamEditMemberArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Name      string   `codec:"name" json:"name"`
	Username  string   `codec:"username" json:"username"`
	Role      TeamRole `codec:"role" json:"role"`
}

func (o TeamEditMemberArg) DeepCopy() TeamEditMemberArg {
	return TeamEditMemberArg{
		SessionID: o.SessionID,
		Name:      o.Name,
		Username:  o.Username,
		Role:      o.Role.DeepCopy(),
	}
}

type TeamsInterface interface {
	TeamCreate(context.Context, TeamCreateArg) error
	TeamGet(context.Context, TeamGetArg) (TeamMembers, error)
	TeamChangeMembership(context.Context, TeamChangeMembershipArg) error
	TeamAddMember(context.Context, TeamAddMemberArg) error
	TeamRemoveMember(context.Context, TeamRemoveMemberArg) error
	TeamEditMember(context.Context, TeamEditMemberArg) error
}

func TeamsProtocol(i TeamsInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.teams",
		Methods: map[string]rpc.ServeHandlerDescription{
			"teamCreate": {
				MakeArg: func() interface{} {
					ret := make([]TeamCreateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamCreateArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamCreateArg)(nil), args)
						return
					}
					err = i.TeamCreate(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"teamGet": {
				MakeArg: func() interface{} {
					ret := make([]TeamGetArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamGetArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamGetArg)(nil), args)
						return
					}
					ret, err = i.TeamGet(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"teamChangeMembership": {
				MakeArg: func() interface{} {
					ret := make([]TeamChangeMembershipArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamChangeMembershipArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamChangeMembershipArg)(nil), args)
						return
					}
					err = i.TeamChangeMembership(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"teamAddMember": {
				MakeArg: func() interface{} {
					ret := make([]TeamAddMemberArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamAddMemberArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamAddMemberArg)(nil), args)
						return
					}
					err = i.TeamAddMember(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"teamRemoveMember": {
				MakeArg: func() interface{} {
					ret := make([]TeamRemoveMemberArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamRemoveMemberArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamRemoveMemberArg)(nil), args)
						return
					}
					err = i.TeamRemoveMember(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"teamEditMember": {
				MakeArg: func() interface{} {
					ret := make([]TeamEditMemberArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamEditMemberArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamEditMemberArg)(nil), args)
						return
					}
					err = i.TeamEditMember(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type TeamsClient struct {
	Cli rpc.GenericClient
}

func (c TeamsClient) TeamCreate(ctx context.Context, __arg TeamCreateArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamCreate", []interface{}{__arg}, nil)
	return
}

func (c TeamsClient) TeamGet(ctx context.Context, __arg TeamGetArg) (res TeamMembers, err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamGet", []interface{}{__arg}, &res)
	return
}

func (c TeamsClient) TeamChangeMembership(ctx context.Context, __arg TeamChangeMembershipArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamChangeMembership", []interface{}{__arg}, nil)
	return
}

func (c TeamsClient) TeamAddMember(ctx context.Context, __arg TeamAddMemberArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamAddMember", []interface{}{__arg}, nil)
	return
}

func (c TeamsClient) TeamRemoveMember(ctx context.Context, __arg TeamRemoveMemberArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamRemoveMember", []interface{}{__arg}, nil)
	return
}

func (c TeamsClient) TeamEditMember(ctx context.Context, __arg TeamEditMemberArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamEditMember", []interface{}{__arg}, nil)
	return
}
