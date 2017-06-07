// Auto-generated by avdl-compiler v1.3.17 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/tlf.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type CryptKeysArg struct {
	Query TLFQuery `codec:"query" json:"query"`
}

func (o CryptKeysArg) DeepCopy() CryptKeysArg {
	return CryptKeysArg{
		Query: o.Query.DeepCopy(),
	}
}

type PublicCanonicalTLFNameAndIDArg struct {
	Query TLFQuery `codec:"query" json:"query"`
}

func (o PublicCanonicalTLFNameAndIDArg) DeepCopy() PublicCanonicalTLFNameAndIDArg {
	return PublicCanonicalTLFNameAndIDArg{
		Query: o.Query.DeepCopy(),
	}
}

type CompleteAndCanonicalizePrivateTlfNameArg struct {
	Query TLFQuery `codec:"query" json:"query"`
}

func (o CompleteAndCanonicalizePrivateTlfNameArg) DeepCopy() CompleteAndCanonicalizePrivateTlfNameArg {
	return CompleteAndCanonicalizePrivateTlfNameArg{
		Query: o.Query.DeepCopy(),
	}
}

type TlfInterface interface {
	// CryptKeys returns TLF crypt keys from all generations.
	CryptKeys(context.Context, TLFQuery) (GetTLFCryptKeysRes, error)
	// * tlfCanonicalID returns the canonical name and TLFID for tlfName.
	// * TLFID should not be cached or stored persistently.
	PublicCanonicalTLFNameAndID(context.Context, TLFQuery) (CanonicalTLFNameAndIDWithBreaks, error)
	CompleteAndCanonicalizePrivateTlfName(context.Context, TLFQuery) (CanonicalTLFNameAndIDWithBreaks, error)
}

func TlfProtocol(i TlfInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.tlf",
		Methods: map[string]rpc.ServeHandlerDescription{
			"CryptKeys": {
				MakeArg: func() interface{} {
					ret := make([]CryptKeysArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]CryptKeysArg)
					if !ok {
						err = rpc.NewTypeError((*[]CryptKeysArg)(nil), args)
						return
					}
					ret, err = i.CryptKeys(ctx, (*typedArgs)[0].Query)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"publicCanonicalTLFNameAndID": {
				MakeArg: func() interface{} {
					ret := make([]PublicCanonicalTLFNameAndIDArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PublicCanonicalTLFNameAndIDArg)
					if !ok {
						err = rpc.NewTypeError((*[]PublicCanonicalTLFNameAndIDArg)(nil), args)
						return
					}
					ret, err = i.PublicCanonicalTLFNameAndID(ctx, (*typedArgs)[0].Query)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"completeAndCanonicalizePrivateTlfName": {
				MakeArg: func() interface{} {
					ret := make([]CompleteAndCanonicalizePrivateTlfNameArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]CompleteAndCanonicalizePrivateTlfNameArg)
					if !ok {
						err = rpc.NewTypeError((*[]CompleteAndCanonicalizePrivateTlfNameArg)(nil), args)
						return
					}
					ret, err = i.CompleteAndCanonicalizePrivateTlfName(ctx, (*typedArgs)[0].Query)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type TlfClient struct {
	Cli rpc.GenericClient
}

// CryptKeys returns TLF crypt keys from all generations.
func (c TlfClient) CryptKeys(ctx context.Context, query TLFQuery) (res GetTLFCryptKeysRes, err error) {
	__arg := CryptKeysArg{Query: query}
	err = c.Cli.Call(ctx, "keybase.1.tlf.CryptKeys", []interface{}{__arg}, &res)
	return
}

// * tlfCanonicalID returns the canonical name and TLFID for tlfName.
// * TLFID should not be cached or stored persistently.
func (c TlfClient) PublicCanonicalTLFNameAndID(ctx context.Context, query TLFQuery) (res CanonicalTLFNameAndIDWithBreaks, err error) {
	__arg := PublicCanonicalTLFNameAndIDArg{Query: query}
	err = c.Cli.Call(ctx, "keybase.1.tlf.publicCanonicalTLFNameAndID", []interface{}{__arg}, &res)
	return
}

func (c TlfClient) CompleteAndCanonicalizePrivateTlfName(ctx context.Context, query TLFQuery) (res CanonicalTLFNameAndIDWithBreaks, err error) {
	__arg := CompleteAndCanonicalizePrivateTlfNameArg{Query: query}
	err = c.Cli.Call(ctx, "keybase.1.tlf.completeAndCanonicalizePrivateTlfName", []interface{}{__arg}, &res)
	return
}
