

## API Description

### Vault

Creates an Ansible vault from string.

**URL** : `/api/v1/vault`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "text": "[string]"
}
```

**Data example**

```json
{
    "text": "some plaintext for encryption"
}
```
**Request example**
```shell
curl --request POST \
  --url http://localhost:8080/api/v1/vault \
  --header 'Content-Type: application/json' \
  --data '{
"text": "text"
}
'
```

#### Success Response

**Code** : `201 Created`

**Content example**

```json
{
    "secret": "$ANSIBLE_VAULT;1.1;AES256\n33613434373839326261633662616431653331663832306362316531643066313163646635393331\n6163393936303663386234633130376464306166643265330a383532633164613664663434616431\n63333263653437616639366235313034376433313133363133306663336136353239376536393665\n3463616234616666320a666334303733376131323831356330633538616534353263353333643365\n3633"
}
```

#### Error Response

**Condition** : If 'text' is not present or misspelled.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "message": "invalid json body",
    "status": 400,
    "error": "bad_request"
}
```
