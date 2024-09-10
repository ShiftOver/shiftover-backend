module.exports = {
  async up(db, client) {
    // users collection
    await db.createCollection('users', {
      validator: {
        $jsonSchema: {
          bsonType: 'object',
          required: ['userId', 'nurseId', 'firstName', 'lastName', 'email', 'profilePictureUrl', 'wardId', 'startDate', 'dateOfBirth', 'position', 'contact', 'createdAt', 'updatedAt'],
          properties: {
            userId: { bsonType: 'string' },
            nurseId: { bsonType: 'string' },
            firstName: { bsonType: 'string' },
            lastName: { bsonType: 'string' },
            email: { bsonType: 'string' },
            profilePictureUrl: { bsonType: 'string' },
            wardId: { bsonType: 'string' },
            startDate: { bsonType: 'string' },
            dateOfBirth: { bsonType: 'string' },
            position: { bsonType: 'string' },
            contact: { bsonType: 'string' },
            createdAt: { bsonType: 'date' },
            updatedAt: { bsonType: 'date' },
          }
        }
      }
    });

    // patients collection
    await db.createCollection('patients', {
      validator: {
        $jsonSchema: {
          bsonType: 'object',
          required: ['patientId', 'firstName', 'lastName', 'dateOfBirth', 'hn', 'gender', 'createdAt', 'updatedAt'],
          properties: {
            patientId: { bsonType: 'string' },
            firstName: { bsonType: 'string' },
            lastName: { bsonType: 'string' },
            dateOfBirth: { bsonType: 'string' },
            hn: { bsonType: 'string' },
            gender: { enum: ["male", "female", "others"] },
            allergies: {
              bsonType: 'array',
              items: {
                bsonType: 'object',
                properties: {
                  name: { bsonType: 'string' },
                },
              },
            },
            createdAt: { bsonType: 'date' },
            updatedAt: { bsonType: 'date' },
          }
        }
      }
    });

    // hospitals collection
    await db.createCollection('hospitals', {
      validator: {
        $jsonSchema: {
          bsonType: 'object',
          required: ['hospitalId', 'hospitalName', 'createdAt'],
          properties: {
            hospitalId: { bsonType: 'string' },
            hospitalName: { bsonType: 'string' },
            createdAt: { bsonType: 'date' },
          }
        }
      }
    });

    // wards collection
    await db.createCollection('wards', {
      validator: {
        $jsonSchema: {
          bsonType: 'object',
          required: ['wardId', 'wardName', 'hospitalId', 'createdAt'],
          properties: {
            wardId: { bsonType: 'string' },
            wardName: { bsonType: 'string' },
            hospitalId: { bsonType: 'string' },
            createdAt: { bsonType: 'date' },
          }
        }
      }
    });

    // rooms collection
    await db.createCollection('rooms', {
      validator: {
        $jsonSchema: {
          bsonType: 'object',
          required: ['roomId', 'roomName', 'wardId', 'createdAt'],
          properties: {
            roomId: { bsonType: 'string' },
            roomName: { bsonType: 'string' },
            wardId: { bsonType: 'string' },
            currentPatientId: { bsonType: 'string' },
            createdAt: { bsonType: 'date' },
          }
        }
      }
    });

    // counters collection
    await db.createCollection('counters');
    await db.collection('counters').insertOne({ _id: 'userId', sequence_value: 0 });

  },

  async down(db, client) {
    await db.collection('users').drop();
    await db.collection('patients').drop();
    await db.collection('hospitals').drop();
    await db.collection('wards').drop();
    await db.collection('rooms').drop();
    await db.collection('counter').drop();
  }
};
