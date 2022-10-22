package util

var payloadSair = `
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
                                    "body": "sair"
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
										"id": "agent_test",
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
                                    "body": "mensagem do Renato para o agent_test"
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

var payloadStatus1 = `
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
								"id": "wamid.HBgNNTUxNjk5MzI1OTI1NhUCABEYEkZDNjI1QjM0NENGRjcxMTc3MwA=",
								"status": "sent",
								"timestamp": "1665616406",
								"recipient_id": "5516993259256",
								"conversation": {
									"id": "78dd8243b1d5e57fbc23883007f4e497",
									"expiration_timestamp": "1665691380",
									"origin": {
										"type": "user_initiated"
									}
								},
								"pricing": {
									"billable": true,
									"pricing_model": "CBP",
									"category": "user_initiated"
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

var payloadStatus2 = `
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
								"id": "wamid.HBgNNTUxNjk5MzI1OTI1NhUCABEYEkZDNjI1QjM0NENGRjcxMTc3MwA=",
								"status": "delivered",
								"timestamp": "1665616407",
								"recipient_id": "5516993259256",
								"conversation": {
									"id": "78dd8243b1d5e57fbc23883007f4e497",
									"origin": {
										"type": "user_initiated"
									}
								},
								"pricing": {
									"billable": true,
									"pricing_model": "CBP",
									"category": "user_initiated"
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

var invalidPayload = `
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
								"id": "wamid.HBgNNTUxNjk5MzI1OTI1NhUCABEYEkZDNjI1QjM0NENGRjcxMTc3MwA=",
								"status": "delivered",
								"timestamp": "1665616407",
								"conversation": {
									"id": "78dd8243b1d5e57fbc23883007f4e497",
									"origin": {
										"type": "user_initiated"
									}
								},
								"pricing": {
									"billable": true,
									"pricing_model": "CBP",
									"category": "user_initiated"
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

var errorPayload = `
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
                                "id": "wamid.HBgNNTUxNjk5MzI1OTI1NhUCABEYEjM3MTI5MTM4MzMxQUU3MEVDQQA=",
                                "status": "failed",
                                "timestamp": "1665854858",
                                "recipient_id": "5516993259256",
                                "errors": [
                                    {
                                        "code": 131047,
                                        "title": "Message failed to send because more than 24 hours have passed since the customer last replied to this number.",
                                        "href": "https:\/\/developers.facebook.com\/docs\/whatsapp\/cloud-api\/support\/error-codes\/"
                                    }
                                ]
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
