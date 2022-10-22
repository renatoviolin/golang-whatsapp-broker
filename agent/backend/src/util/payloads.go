package util

var payload1 = `
{
    "object": "whatsapp_business_account",
    "entry": [
        {
            "id": "103471172512719",
            "changes": [
                {
                    "value": {
                        "messaging_product": "whatsapp",
                        "metadata": {
                            "display_phone_number": "15550915609",
                            "phone_number_id": "102261592636695"
                        },
                        "contacts": [
                            {
                                "profile": {
                                    "name": "Renato"
                                },
                                "wa_id": "5516993259256"
                            }
                        ],
                        "messages": [
                            {
                                "from": "5516993259256",
                                "id": "wamid.HBgNNTUxNjk5MzI1OTI1NhUCABIYFDNBMTdFMTAzODIwMUEwMDIyOTJBAA==",
                                "timestamp": "1663444803",
                                "text": {
                                    "body": "olá, gostaria de ser atendido"
                                },
                                "type": "text"
                            }
                        ]
                    },
                    "field": "messages"
                }
            ]
        }
    ]
}
`

var payload2 = `
{
	"object": "whatsapp_business_account",
	"entry": [
		{
			"id": "103471172512719",
			"changes": [
				{
					"value": {
						"messaging_product": "whatsapp",
						"metadata": {
							"display_phone_number": "15550915609",
							"phone_number_id": "102261592636695"
						},
						"contacts": [
							{
								"profile": {
									"name": "Renato"
								},
								"wa_id": "5516993259256"
							}
						],
						"messages": [
							{
								"context": {
									"from": "15550915609",
									"id": "wamid.HBgNNTUxNjk5MzI1OTI1NhUCABEYEkZFQzg2NEMxQzAwQTE0MjAyQQA="
								},
								"from": "5516993259256",
								"id": "wamid.HBgNNTUxNjk5MzI1OTI1NhUCABIYFDNFQjBENTIwRDk5NzZBMkVBNzEwAA==",
								"timestamp": "1664232480",
								"type": "interactive",
								"interactive": {
									"type": "list_reply",
									"list_reply": {
										"id": "agent_a",
										"title": "Agente A"
									}
								}
							}
						]
					},
					"field": "messages"
				}
			]
		}
	]
}
`

var payloadInvalidAgentID = `
{
	"object": "whatsapp_business_account",
	"entry": [
		{
			"id": "103471172512719",
			"changes": [
				{
					"value": {
						"messaging_product": "whatsapp",
						"metadata": {
							"display_phone_number": "15550915609",
							"phone_number_id": "102261592636695"
						},
						"contacts": [
							{
								"profile": {
									"name": "Renato"
								},
								"wa_id": "5516993259256"
							}
						],
						"messages": [
							{
								"context": {
									"from": "15550915609",
									"id": "wamid.HBgNNTUxNjk5MzI1OTI1NhUCABEYEkZFQzg2NEMxQzAwQTE0MjAyQQA="
								},
								"from": "5516993259256",
								"id": "wamid.HBgNNTUxNjk5MzI1OTI1NhUCABIYFDNFQjBENTIwRDk5NzZBMkVBNzEwAA==",
								"timestamp": "1664232480",
								"type": "interactive",
								"interactive": {
									"type": "list_reply",
									"list_reply": {
										"id": "agent_xxx",
										"title": "Agente X"
									}
								}
							}
						]
					},
					"field": "messages"
				}
			]
		}
	]
}
`

var payload3 = `
{
    "object": "whatsapp_business_account",
    "entry": [
        {
            "id": "103471172512719",
            "changes": [
                {
                    "value": {
                        "messaging_product": "whatsapp",
                        "metadata": {
                            "display_phone_number": "15550915609",
                            "phone_number_id": "102261592636695"
                        },
                        "contacts": [
                            {
                                "profile": {
                                    "name": "Renato"
                                },
                                "wa_id": "5516993259256"
                            }
                        ],
                        "messages": [
                            {
                                "from": "5516993259256",
                                "id": "wamid.HBgNNTUxNjk5MzI1OTI1NhUCABIYFDNBMTdFMTAzODIwMUEwMDIyOTJBAA==",
                                "timestamp": "1663444803",
                                "text": {
                                    "body": "mensagem do Renato para o agent_a"
                                },
                                "type": "text"
                            }
                        ]
                    },
                    "field": "messages"
                }
            ]
        }
    ]
}
`

var payloadResponseButton = `
{
	"object":"whatsapp_business_account",
	"entry":[
	   {
		  "id":"103471172512719",
		  "changes":[
			 {
				"value":{
				   "messaging_product":"whatsapp",
				   "metadata":{
					  "display_phone_number":"15550915609",
					  "phone_number_id":"102261592636695"
				   },
				   "contacts":[
					  {
						 "profile":{
							"name":"Renato"
						 },
						 "wa_id":"5516993259256"
					  }
				   ],
				   "messages":[
					  {
						 "context":{
							"from":"15550915609",
							"id":"wamid.HBgNNTUxNjk5MzI1OTI1NhUCABEYEjAxNEVDNDc3RUQxNTQxQ0JFQQA="
						 },
						 "from":"5516993259256",
						 "id":"wamid.HBgNNTUxNjk5MzI1OTI1NhUCABIYFDNBQzAxQjlERjU3NDE2QzEwMjA3AA==",
						 "timestamp":"1663759701",
						 "type":"interactive",
						 "interactive":{
							"type":"list_reply",
							"list_reply":{
							   "id":"id1",
							   "title":"row-title-1",
							   "description":"row-description-1"
							}
						 }
					  }
				   ]
				},
				"field":"messages"
			 }
		  ]
	   }
	]
 }
`

var payloadInvalid = `
{
    "object": "whatsapp_business_account",
    "entry": [
        {
            "id": "103471172512719",
            "changes": [
                {
                    "value": {
                        "messaging_product": "whatsapp",
                        "metadata": {
                            "display_phone_number": "15550915609",
                            "phone_number_id": "102261592636695"
                        }
                    },
                    "field": "messages"
                }
            ]
        }
    ]
}
`

var payloadStatus = `
{
    "object": "whatsapp_business_account",
    "entry": [
        {
            "id": "103471172512719",
            "changes": [
                {
                    "value": {
                        "messaging_product": "whatsapp",
                        "metadata": {
                            "display_phone_number": "15550915609",
                            "phone_number_id": "102261592636695"
                        },
                        "statuses": [
                            {
                                "id": "wamid.HBgNNTUxNjk5MzI1OTI1NhUCABEYEjkwNTZFRkIyQkRFRUYwQTgxNgA=",
                                "status": "delivered",
                                "timestamp": "1663449445",
                                "recipient_id": "111",
                                "conversation": {
                                    "id": "c7662d4fa5d580e5238260cdbb77483a",
                                    "origin": {
                                        "type": "business_initiated"
                                    }
                                },
                                "pricing": {
                                    "billable": true,
                                    "pricing_model": "CBP",
                                    "category": "business_initiated"
                                }
                            }
                        ]
                    },
                    "field": "messages"
                }
            ]
        }
    ]
}
`
