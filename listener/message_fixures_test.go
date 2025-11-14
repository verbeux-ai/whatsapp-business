package listener_test

const textMessage = `{
    "entry": [
      {
        "changes": [
          {
            "field": "messages",
            "value": {
              "contacts": [
                {
                  "profile": {
                    "name": "Pedro Ivo"
                  },
                  "wa_id": "558598437440"
                }
              ],
              "messages": [
                {
                  "from": "558598437440",
                  "id": "wamid.HBgMNTU4NTk4NDM3NDQwFQIAEhgUM0EzOTU5OUZENEMxMDQyNkE2QUEA",
                  "text": {
                    "body": "teste"
                  },
                  "timestamp": "1746579317",
                  "type": "text"
                }
              ],
              "messaging_product": "whatsapp",
              "metadata": {
                "display_phone_number": "15550077211",
                "phone_number_id": "245796831949065"
              }
            }
          }
        ],
        "id": "190897907451133"
      }
    ],
    "object": "whatsapp_business_account"
  }`

const audioMessage = `{
    "entry": [
      {
        "changes": [
          {
            "field": "messages",
            "value": {
              "contacts": [
                {
                  "profile": {
                    "name": "Pedro Ivo"
                  },
                  "wa_id": "558598437440"
                }
              ],
              "messages": [
                {
                  "audio": {
                    "id": "1246869980292947",
                    "mime_type": "audio/ogg; codecs=opus",
                    "sha256": "gbabQ3akK+Pv8IiA744G8cdW5RGdex/rXkpfuM+VBBc=",
                    "voice": true
                  },
                  "from": "558598437440",
                  "id": "wamid.HBgMNTU4NTk4NDM3NDQwFQIAEhgUM0FDRDRGMUM1QUU2OEQ1QkY1OTAA",
                  "timestamp": "1746578663",
                  "type": "audio"
                }
              ],
              "messaging_product": "whatsapp",
              "metadata": {
                "display_phone_number": "15550077211",
                "phone_number_id": "245796831949065"
              }
            }
          }
        ],
        "id": "190897907451133"
      }
    ],
    "object": "whatsapp_business_account"
  }`

const imageMessage = `{
    "entry": [
      {
        "changes": [
          {
            "field": "messages",
            "value": {
              "contacts": [
                {
                  "profile": {
                    "name": "Pedro Ivo"
                  },
                  "wa_id": "558598437440"
                }
              ],
              "messages": [
                {
                  "from": "558598437440",
                  "id": "wamid.HBgMNTU4NTk4NDM3NDQwFQIAEhgUM0E5QjlFQjA2NEE0QjE5OTEzRjcA",
                  "image": {
                    "caption": "asdasdasd",
                    "id": "1930575551090839",
                    "mime_type": "image/jpeg",
                    "sha256": "U3iYGwGrXsj1BlguLhseaLeLJEg9n0KlGxwVfkHqEYI="
                  },
                  "timestamp": "1746580151",
                  "type": "image"
                }
              ],
              "messaging_product": "whatsapp",
              "metadata": {
                "display_phone_number": "15550077211",
                "phone_number_id": "245796831949065"
              }
            }
          }
        ],
        "id": "190897907451133"
      }
    ],
    "object": "whatsapp_business_account"
  }`

const documentMessage = `{
    "entry": [
      {
        "changes": [
          {
            "field": "messages",
            "value": {
              "contacts": [
                {
                  "profile": {
                    "name": "Pedro Ivo"
                  },
                  "wa_id": "558598437440"
                }
              ],
              "messages": [
                {
                  "document": {
                    "caption": "aaaaaaaaaaa",
                    "filename": "Trancamento_de_Matricula.pdf",
                    "id": "1233520308187160",
                    "mime_type": "application/pdf",
                    "sha256": "IWJCna4CfvFsjGUD9JrFdjdjW1jjyf+CPu200+yQn6k="
                  },
                  "from": "558598437440",
                  "id": "wamid.HBgMNTU4NTk4NDM3NDQwFQIAEhgUM0E3MUQzMTZFQTc3NzBBRDU3MDEA",
                  "timestamp": "1746580635",
                  "type": "document"
                }
              ],
              "messaging_product": "whatsapp",
              "metadata": {
                "display_phone_number": "15550077211",
                "phone_number_id": "245796831949065"
              }
            }
          }
        ],
        "id": "190897907451133"
      }
    ],
    "object": "whatsapp_business_account"
  }`

const statusesMessage = `{
  "object": "whatsapp_business_account",
  "entry": [
    {
      "id": "190897907451133",
      "changes": [
        {
          "value": {
            "messaging_product": "whatsapp",
            "metadata": {
              "display_phone_number": "15550077211",
              "phone_number_id": "245796831949065"
            },
            "statuses": [
              {
                "id": "wamid.HBgMNTU4NTk4NDM3NDQwFQIAERgSMUU4NzE3QkFGMjkyQjk5RTJGAA==",
                "status": "delivered",
                "timestamp": "1746590336",
                "recipient_id": "558598437440",
                "conversation": {
                  "id": "6852cd1442983d45cbdc8958f5f3f066",
                  "origin": {
                    "type": "utility"
                  }
                },
                "pricing": {
                  "billable": true,
                  "pricing_model": "CBP",
                  "category": "utility"
                }
              }
            ]
          },
          "field": "messages"
        }
      ]
    }
  ]
}`

const quickReplyMessage = `{
  "object" : "whatsapp_business_account",
  "entry" : [ {
    "id" : "190897907451133",
    "changes" : [ {
      "value" : {
        "messaging_product" : "whatsapp",
        "metadata" : {
          "display_phone_number" : "15550077211",
          "phone_number_id" : "245796831949065"
        },
        "contacts" : [ {
          "profile" : {
            "name" : "Ivo"
          },
          "wa_id" : "558594138387"
        } ],
        "messages" : [ {
          "context" : {
            "from" : "15550077211",
            "id" : "wamid.HBgMNTU4NTk0MTM4Mzg3FQIAERgSMzQ1RkZGNEI4MjhDNjNGNDZEAA=="
          },
          "from" : "558594138387",
          "id" : "wamid.HBgMNTU4NTk0MTM4Mzg3FQIAEhgWM0VCMDI3MEYyQ0JGOUQ0NEZCOUU5MQA=",
          "timestamp" : "1761312711",
          "type" : "button",
          "button" : {
            "payload" : "Nao testei",
            "text" : "Nao testei"
          }
        } ]
      },
      "field" : "messages"
    } ]
  } ]
}`

const listMessage = `{"object":"whatsapp_business_account","entry":[{"id":"190897907451133","changes":[{"value":{"messaging_product":"whatsapp","metadata":{"display_phone_number":"15550077211","phone_number_id":"245796831949065"},"contacts":[{"profile":{"name":"Ivo"},"wa_id":"558594138387"}],"messages":[{"context":{"from":"15550077211","id":"wamid.HBgMNTU4NTk0MTM4Mzg3FQIAERgSRUEzNzVGMTM3RjIyMDBEQjAyAA=="},"from":"558594138387","id":"wamid.HBgMNTU4NTk0MTM4Mzg3FQIAEhgWM0VCMDZFRkVCOEQyOUYzRUI2RTcwNwA=","timestamp":"1763120309","type":"interactive","interactive":{"type":"list_reply","list_reply":{"id":"2","title":"Editar informa\u00e7\u00f5es","description":"Editar informa\u00e7\u00f5es"}}}]},"field":"messages"}]}]}`
