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
