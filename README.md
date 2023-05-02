# happiness-to-straycat

## List of Packages, Features and Database Used

- SQLC 
- Fiber V2
- UUID
- JWT
- Docker
- Viper
- Postgresql
- Redis
- Compile Daemon

URL to Test https://happiness-to-straycat-production.up.railway.app/

## USER 

register> api/auth/register [POST]

>{
  "email": "user@example.com",
  "password": "password123",
  "name": "John Doe",
  "gender": "male",
  "age": 30,
  "address": "123 Main St",
  "phone": 5551234567
}

Login> api/auth/login [POST]

>{
  "email":"user@example.com",
  "password":"password123"
}

Check Login User> /api/users/me [GET]

Logout> api/auth/logout [GET]

## Fav_Food

Create Tag> api/food/ [POST]

>{
    "Company": "NewFood",
    "Variety": "Test",
    "Protein": 6000,
    "Fat": 16150,
    "Carbs": 65450,
    "Phos": 554640,
    "Notes": "Test"
}

Update Tag> api/food/1 [PATCH]

Get TagByID> api/food/1 [GET]

Get All Tag> api/food/ [GET]

Delete TagByID> api/food/1 [DELETE]

## Immunization

Create Immunization> api/immunization/ [POST]

>{
  "name":"New Immunization "
}

Update Immunization> api/immunization/1 [PATCH]

Get ImmunizationByID> api/immunization/1 [GET]

Get All Immunization> api/immunization/ [GET]

Delete ImmunizationByID> api/immunization/1 [DELETE]

## Tag

Create Tag> api/tag/ [POST]

>{
  "name":"New TagName"
}


Update Tag> api/tag/1 [PATCH]

Get TagByID> api/tag/1 [GET]

Get All Tag> api/tag/ [GET]

Delete TagByID> api/tag/1 [DELETE]
