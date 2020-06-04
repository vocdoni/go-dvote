package vochain

// Testnet Genesis File for Vocdoni KISS v1
const (
	ReleaseGenesis1 = `
{
  "genesis_time": "2020-06-03T09:30:50.199392102Z",
  "chain_id": "vocdoni-oc2020-02",
  "consensus_params": {
    "block": {
      "max_bytes": "22020096",
      "max_gas": "-1",
      "time_iota_ms": "10000"
    },
    "evidence": {
      "max_age_num_blocks": "100000",
      "max_age_duration": "10000"
    },
    "validator": {
      "pub_key_types": [
        "ed25519"
      ]
    }
  },
  "validators": [
    {
      "address": "8818C252BD0FD0A06B06CBEA79CA6DF8872996FA",
      "pub_key": {
        "type": "tendermint/PubKeyEd25519",
        "value": "UhAW972a9FDsKmpF9ayh5VdfUo6A0AAWiE9ztleHBqs="
      },
      "power": "10",
      "name": ""
    },
    {
      "address": "40AB065BCCC03EC8A7160E1FC125C6F17BA9FAE0",
      "pub_key": {
        "type": "tendermint/PubKeyEd25519",
        "value": "4Quwe8cmefUiLND8Zq4bz2KHdpzpu+RNxez/w/mGGzA="
      },
      "power": "10",
      "name": ""
    },
    {
      "address": "5BF6C1C067FC61C4C09CC3D930D27DC411210902",
      "pub_key": {
        "type": "tendermint/PubKeyEd25519",
        "value": "RIZi6LSQDrTA+5kRFKqP08SbG7GB+YR03LwprokPkx4="
      },
      "power": "10",
      "name": ""
    },
    {
      "address": "F3BEA2B40DF378CC1FFB621546629C82CE21EFC9",
      "pub_key": {
        "type": "tendermint/PubKeyEd25519",
        "value": "bZIr9Sn8vLpgnNPI0xmYNzjLBud4MBHTR17BfR5z+xc="
      },
      "power": "10",
      "name": ""
    }
  ],
  "app_hash": "",
  "app_state": {
    "validators": [
      {
        "address": "8818C252BD0FD0A06B06CBEA79CA6DF8872996FA",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "UhAW972a9FDsKmpF9ayh5VdfUo6A0AAWiE9ztleHBqs="
        },
        "power": "10",
        "name": "miner1"
      },
      {
        "address": "40AB065BCCC03EC8A7160E1FC125C6F17BA9FAE0",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "4Quwe8cmefUiLND8Zq4bz2KHdpzpu+RNxez/w/mGGzA="
        },
        "power": "10",
        "name": "miner2"
      },
      {
        "address": "5BF6C1C067FC61C4C09CC3D930D27DC411210902",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "RIZi6LSQDrTA+5kRFKqP08SbG7GB+YR03LwprokPkx4="
        },
        "power": "10",
        "name": "miner3"
      },
      {
        "address": "F3BEA2B40DF378CC1FFB621546629C82CE21EFC9",
        "pub_key": {
          "type": "tendermint/PubKeyEd25519",
          "value": "bZIr9Sn8vLpgnNPI0xmYNzjLBud4MBHTR17BfR5z+xc="
        },
        "power": "10",
        "name": "miner4"
      }
        ],
    "oracles": [
      "ba8bf885892506f7c957f1623e144ee1a25a66c7",
	  "dab04e20f00d7967fc0e7dc99addf6194e1b5af7"
    ]
  }
}
  `

	DevelopmentGenesis1 = `
{
   "genesis_time":"2020-06-01T20:29:50.512370579Z",
   "chain_id":"vocdoni-development-13",
   "consensus_params":{
      "block":{
         "max_bytes":"22020096",
         "max_gas":"-1",
         "time_iota_ms":"10000"
      },
      "evidence":{
         "max_age_num_blocks":"100000",
         "max_age_duration":"10000"
      },
      "validator":{
         "pub_key_types":[
            "ed25519"
         ]
      }
   },
   "validators":[
      {
         "address":"5C69093136E0CB84E5CFA8E958DADB33C0D0CCCF",
         "pub_key":{
            "type":"tendermint/PubKeyEd25519",
            "value":"mXc5xXTKgDSYcy1lBCT1Ag7Lh1nPWHMa/p80XZPzAPY="
         },
         "power":"10",
         "name":"miner0"
      },
      {
         "address":"2E1B244B84E223747126EF621C022D5CEFC56F69",
         "pub_key":{
            "type":"tendermint/PubKeyEd25519",
            "value":"gaf2ZfdxpoielRXDXyBcMxkdzywcE10WsvLMe1K62UY="
         },
         "power":"10",
         "name":"miner1"
      },
      {
         "address":"4EF00A8C18BD472167E67F28694F31451A195581",
         "pub_key":{
            "type":"tendermint/PubKeyEd25519",
            "value":"dZXMBiQl4s0/YplfX9iMnCWonJp2gjrFHHXaIwqqtmc="
         },
         "power":"10",
         "name":"miner2"
      },
      {
         "address":"ECCC09A0DF8F4E5554A9C58F634E9D6AFD5F1598",
         "pub_key":{
            "type":"tendermint/PubKeyEd25519",
            "value":"BebelLYe4GZKwy9IuXCyBTySxQCNRrRoi1DSvAf6QxE="
         },
         "power":"10",
         "name":"miner3"
      },
      {
         "address":"3272B3046C31D87F92E26D249B97CC144D835DA6",
         "pub_key":{
            "type":"tendermint/PubKeyEd25519",
            "value":"hUz8jCePBfG4Bi9s13IdleWq5MZ5upe03M+BX2ah7c4="
         },
         "power":"10",
         "name":"miner4"
      },
      {
         "address":"05BA8FCBEA4A4EDCFD49081B42CA3F9ED13246C1",
         "pub_key":{
            "type":"tendermint/PubKeyEd25519",
            "value":"1FGNernnvg4QpV7psYFQPeIFZJm32yN1SjZULbliidg="
         },
         "power":"10",
         "name":"miner5"
      }
   ],
   "app_hash":"",
   "app_state":{
      "validators":[
         {
            "address":"5C69093136E0CB84E5CFA8E958DADB33C0D0CCCF",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"mXc5xXTKgDSYcy1lBCT1Ag7Lh1nPWHMa/p80XZPzAPY="
            },
            "power":"10",
            "name":"miner0"
         },
         {
            "address":"2E1B244B84E223747126EF621C022D5CEFC56F69",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"gaf2ZfdxpoielRXDXyBcMxkdzywcE10WsvLMe1K62UY="
            },
            "power":"10",
            "name":"miner1"
         },
         {
            "address":"4EF00A8C18BD472167E67F28694F31451A195581",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"dZXMBiQl4s0/YplfX9iMnCWonJp2gjrFHHXaIwqqtmc="
            },
            "power":"10",
            "name":"miner2"
         },
         {
            "address":"ECCC09A0DF8F4E5554A9C58F634E9D6AFD5F1598",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"BebelLYe4GZKwy9IuXCyBTySxQCNRrRoi1DSvAf6QxE="
            },
            "power":"10",
            "name":"miner3"
         },
         {
            "address":"3272B3046C31D87F92E26D249B97CC144D835DA6",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"hUz8jCePBfG4Bi9s13IdleWq5MZ5upe03M+BX2ah7c4="
            },
            "power":"10",
            "name":"miner4"
         },
         {
            "address":"05BA8FCBEA4A4EDCFD49081B42CA3F9ED13246C1",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"1FGNernnvg4QpV7psYFQPeIFZJm32yN1SjZULbliidg="
            },
            "power":"10",
            "name":"miner5"
         }
      ],
      "oracles":[
         "0xb926be24A9ca606B515a835E91298C7cF0f2846f",
         "0x2f4ed2773dcf7ad0ec15eb84ec896f4eebe0e08a"
      ]
   }
}
`
)
