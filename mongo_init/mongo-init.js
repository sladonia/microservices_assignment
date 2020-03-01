db.createUser(
    {
        user: "user",
        pwd: "password",
        roles: [
            {
                role: "readWrite",
                db: "port_db"
            }
        ]
    }
);
db.createCollection("ports");
db.ports.createIndex({ abbreviation: 1 }, { unique: true });
