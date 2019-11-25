// Auto-generated to Go types and interfaces using avdl-compiler v1.4.6 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/gregor1/blocking.avdl

package gregor1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
	"time"
)

type BlockConversationsArg struct {
	Uid    UID      `codec:"uid" json:"uid"`
	TlfIDs []string `codec:"tlfIDs" json:"tlfIDs"`
}

type BlockingInterface interface {
	BlockConversations(context.Context, BlockConversationsArg) error
}

func BlockingProtocol(i BlockingInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "gregor.1.blocking",
		Methods: map[string]rpc.ServeHandlerDescription{
			"blockConversations": {
				MakeArg: func() interface{} {
					var ret [1]BlockConversationsArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]BlockConversationsArg)
					if !ok {
						err = rpc.NewTypeError((*[1]BlockConversationsArg)(nil), args)
						return
					}
					err = i.BlockConversations(ctx, typedArgs[0])
					return
				},
			},
		},
	}
}

type BlockingClient struct {
	Cli rpc.GenericClient
}

func (c BlockingClient) BlockConversations(ctx context.Context, __arg BlockConversationsArg) (err error) {
	err = c.Cli.Call(ctx, "gregor.1.blocking.blockConversations", []interface{}{__arg}, nil, 0*time.Millisecond)
	return
}
