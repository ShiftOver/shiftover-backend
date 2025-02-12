module.exports = {
  async up(db, client) {
    // users collection
    await db.createCollection('users', {
      validator: {
        $jsonSchema: {
          bsonType: 'object',
          required: ['userId', 'nurseId', 'firstName', 'lastName', 'email', 'wardId', 'startDate', 'dateOfBirth', 'position', 'contact', 'createdAt', 'updatedAt'],
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
          required: ['patientId', 'firstName', 'lastName', 'dateOfBirth', 'sex', 'createdAt', 'updatedAt'],
          properties: {
            patientId: { bsonType: 'string' },
            firstName: { bsonType: 'string' },
            lastName: { bsonType: 'string' },
            profilePictureUrl: { bsonType: 'string' },
            sex: { enum: ['male', 'female', 'other'] },
            education: { bsonType: 'string' },
            occupation: { bsonType: 'string' },
            dateOfBirth: { bsonType: 'date' },
            height: { bsonType: 'int' },
            weight: { bsonType: 'int' },
            modeOfArrival: { bsonType: 'string' },
            admittedForm: { bsonType: 'string' },
            initialVitalSigns: {
              bsonType: 'object',
              properties: {
                temperature: { bsonType: 'double' },
                heartRate: { bsonType: 'int' },
                respiratoryRate: { bsonType: 'int' },
                bloodPressure: {
                  bsonType: 'object',
                  properties: {
                    systol: { bsonType: 'int' },
                    diastol: { bsonType: 'int' },
                  },
                },
              },
            },
            diagnosis: { bsonType: 'string' },
            chiefComplaint: { bsonType: 'string' },
            pastIllness: { bsonType: 'string' },
            pastIllnessHistory: { bsonType: 'string' },
            familyIllnessHistory: { bsonType: 'string' },
            allergies: {
              bsonType: 'array',
              items: {
                bsonType: 'object',
                properties: {
                  name: { bsonType: 'string' },
                },
              },
            },
            reactions: {
              bsonType: 'array',
              items: {
                bsonType: 'object',
                properties: {
                  data: { bsonType: 'string' },
                },
              },
            },
            tobacco: {
              bsonType: 'object',
              properties: {
                status: { bsonType: 'string' },
                quitInfo: {
                  bsonType: 'object',
                  properties: {
                    smokedDuration: { bsonType: 'string' },
                    quitDuration: { bsonType: 'string' },
                  },
                },
                continuousInfo: {
                  bsonType: 'object',
                  properties: {
                    duration: { bsonType: 'string' },
                  },
                },
              },
            },
            alcohol: {
              bsonType: 'object',
              properties: {
                status: { bsonType: 'string' },
                quitInfo: {
                  bsonType: 'object',
                  properties: {
                    smokedDuration: { bsonType: 'string' },
                    quitDuration: { bsonType: 'string' },
                  },
                },
                continuousInfo: {
                  bsonType: 'object',
                  properties: {
                    frequency: { bsonType: 'string' },
                    duration: { bsonType: 'string' },
                  },
                },
              },
            },
            drugs: {
              bsonType: 'object',
              properties: {
                status: { bsonType: 'string' },
                quitInfo: {
                  bsonType: 'object',
                  properties: {
                    smokedDuration: { bsonType: 'string' },
                    quitDuration: { bsonType: 'string' },
                  },
                },
                continuousInfo: {
                  bsonType: 'object',
                  properties: {
                    frequency: { bsonType: 'string' },
                    duration: { bsonType: 'string' },
                  },
                },
              },
            },
            exercise: {
              bsonType: 'object',
              properties: {
                status: { bsonType: 'string' },
                frequency: { bsonType: 'string' },
              },
            },
            sleep: {
              bsonType: 'object',
              properties: {
                amount: { bsonType: 'int' },
                status: { bsonType: 'string' },
                helper: { bsonType: 'string' },
              },
            },
            information: {
              bsonType: 'object',
              properties: {
                providedBy: { bsonType: 'string' },
                emergencyContact: {
                  bsonType: 'object',
                  properties: {
                    name: { bsonType: 'string' },
                    relationship: { bsonType: 'string' },
                    phoneNumber: { bsonType: 'string' },
                  },
                },
              },
            },
            createdAt: { bsonType: 'date' },
            updatedAt: { bsonType: 'date' },
          },
        },
      },
    });

    await db.createCollection('patient_assessments', {
      validator: {
        $jsonSchema: {
          bsonType: "object",
          required: ["patientId", "createdAt", "updatedAt"],
          properties: {
            patientId: { bsonType: "string" },
            spiritual: {
              bsonType: "object",
              properties: {
                religion: { bsonType: "string" },
                specialConsideration: { bsonType: "string" },
                anxiety: { bsonType: "string" },
                supportSystem: { bsonType: "string" }
              }
            },
            nutrition: {
              bsonType: "object",
              properties: {
                dietType: { bsonType: "string" },
                appetite: { bsonType: "string" },
                specialDiet: { bsonType: "string" },
                feeding: { bsonType: "string" },
                swallowing: { bsonType: "string" },
                glProblem: { bsonType: "string" },
                weightChange: {
                  bsonType: "object",
                  properties: {
                    status: { bsonType: "string" },
                    change: {
                      bsonType: "object",
                      properties: {
                        type: { bsonType: "string" },
                        weight: { bsonType: "string" },
                        periodValue: { bsonType: "int" },
                        periodType: { bsonType: "string" }
                      }
                    }
                  }
                }
              }
            },
            pulmonary: {
              bsonType: "object",
              properties: {
                rate: { bsonType: "string" },
                rhythm: { bsonType: "string" },
                effort: { bsonType: "string" },
                cough: { bsonType: "string" },
                sputum: { bsonType: "string" }
              }
            },
            cardiovascular: {
              bsonType: "object",
              properties: {
                pulseRhythm: { bsonType: "string" },
                pulseAmplitude: { bsonType: "string" },
                pulseRate: { bsonType: "string" },
                edema: { bsonType: "string" },
                neckVeinEngorged: { bsonType: "string" },
                chestPain: {
                  bsonType: "object",
                  properties: {
                    yes: { bsonType: "bool" },
                    location: { bsonType: "string" },
                    referredPain: { bsonType: "string" },
                    duration: { bsonType: "string" },
                    frequency: { bsonType: "string" }
                  }
                }
              }
            },
            cardioCurrentTreatment: { bsonType: "string" },
            neurosensory: {
              bsonType: "object",
              properties: {
                levelOfConsciousness: { bsonType: "string" },
                vision: {
                  bsonType: "object",
                  properties: {
                    lt: { bsonType: "bool" },
                    rt: { bsonType: "bool" },
                    device: { bsonType: "string" }
                  }
                },
                hearing: {
                  bsonType: "object",
                  properties: {
                    lt: { bsonType: "bool" },
                    rt: { bsonType: "bool" },
                    device: { bsonType: "string" }
                  }
                },
                speech: {
                  bsonType: "object",
                  properties: {
                    normal: { bsonType: "bool" },
                    device: { bsonType: "string" }
                  }
                },
                smell: {
                  bsonType: "object",
                  properties: {
                    normal: { bsonType: "bool" },
                    description: { bsonType: "string" }
                  }
                },
                sensation: { bsonType: "string" }
              }
            },
            musculoskeletal: {
              bsonType: "object",
              properties: {
                handGrasps: {
                  bsonType: "object",
                  properties: {
                    status: { bsonType: "string" },
                    rt: { bsonType: "string" },
                    lt: { bsonType: "string" }
                  }
                },
                joint: {
                  bsonType: "object",
                  properties: {
                    status: { bsonType: "string" },
                    description: { bsonType: "string" }
                  }
                },
                weakness: {
                  bsonType: "object",
                  properties: {
                    yes: { bsonType: "bool" },
                    description: { bsonType: "string" }
                  }
                },
                paralysis: {
                  bsonType: "object",
                  properties: {
                    yes: { bsonType: "bool" },
                    description: { bsonType: "string" }
                  }
                },
                seizure: {
                  bsonType: "object",
                  properties: {
                    yes: { bsonType: "bool" },
                    description: { bsonType: "string" }
                  }
                },
                movement: {
                  bsonType: "object",
                  properties: {
                    yes: { bsonType: "bool" },
                    description: { bsonType: "string" }
                  }
                }
              }
            },
            mobility: {
              bsonType: "object",
              properties: {
                activity: {
                  bsonType: "object",
                  properties: {
                    independent: { bsonType: "bool" },
                    assistance: {
                      bsonType: ["array", "null"],
                      items: { bsonType: "string" }
                    }
                  }
                },
                prostheticDevices: {
                  bsonType: "object",
                  properties: {
                    cane: { bsonType: "bool" },
                    walker: { bsonType: "bool" },
                    wheelChair: { bsonType: "bool" },
                    artificialLimb: { bsonType: "string" },
                    other: { bsonType: "string" }
                  }
                },
                injury: {
                  bsonType: "object",
                  properties: {
                    yes: { bsonType: "bool" },
                    description: { bsonType: "string" }
                  }
                },
                dominantHand: { bsonType: "string" }
              }
            },
            teaching: {
              bsonType: ["array", "null"],
              items: { bsonType: "string" }
            },
            gastrointestinal: {
              bsonType: "object",
              properties: {
                oralCavity: { bsonType: "string" },
                abdomen: { bsonType: "string" },
                bowelPattern: {
                  bsonType: "object",
                  properties: {
                    frequency: { bsonType: "int" },
                    days: { bsonType: "int" }
                  }
                },
                eliminationProblem: { bsonType: "string" }
              }
            },
            genitourinary: {
              bsonType: "object",
              properties: {
                bladder: { bsonType: "string" },
                voiding: {
                  bsonType: "object",
                  properties: {
                    day: { bsonType: "int" },
                    night: { bsonType: "int" },
                    status: { bsonType: "string" }
                  }
                },
                urine: { bsonType: "string" }
              }
            },
            painManagement: {
              bsonType: "object",
              properties: {
                pain: {
                  bsonType: "object",
                  properties: {
                    yes: { bsonType: "bool" },
                    description: { bsonType: "string" },
                    when: { bsonType: "string" },
                    cause: { bsonType: "string" }
                  }
                },
                pattern: { bsonType: "string" },
                patientDescribe: { bsonType: "string" },
                intensity: { bsonType: "int" },
                affectedAbility: { bsonType: "string" },
                relieves: { bsonType: "string" }
              }
            },
            discharge: {
              bsonType: "object",
              properties: {
                screeningCriteria: {
                  bsonType: "object",
                  properties: {
                    postDischargeAssistance: {
                      bsonType: "object",
                      properties: {
                        yes: { bsonType: "bool" },
                        familyCapable: {
                          bsonType: "object",
                          properties: {
                            yes: { bsonType: "bool" },
                            description: { bsonType: "string" }
                          }
                        },
                        familyCannotProvide: {
                          bsonType: "object",
                          properties: {
                            yes: { bsonType: "bool" },
                            description: { bsonType: "string" }
                          }
                        }
                      }
                    },
                    financialConcern: {
                      bsonType: "object",
                      properties: {
                        yes: { bsonType: "bool" },
                        description: { bsonType: "string" }
                      }
                    }
                  }
                },
                homeEnvironment: {
                  bsonType: "object",
                  properties: {
                    liveWith: { bsonType: "string" },
                    liveWhere: { bsonType: "string" }
                  }
                },
                planningNeeds: {
                  bsonType: "object",
                  properties: {
                    medication: {
                      bsonType: "object",
                      properties: {
                        yes: { bsonType: "bool" },
                        description: { bsonType: "string" }
                      }
                    },
                    environment: {
                      bsonType: "object",
                      properties: {
                        yes: { bsonType: "bool" },
                        description: { bsonType: "string" }
                      }
                    },
                    treatment: {
                      bsonType: "object",
                      properties: {
                        yes: { bsonType: "bool" },
                        description: { bsonType: "string" }
                      }
                    },
                    health: {
                      bsonType: "object",
                      properties: {
                        yes: { bsonType: "bool" },
                        description: { bsonType: "string" }
                      }
                    },
                    outpatientReferral: {
                      bsonType: "object",
                      properties: {
                        yes: { bsonType: "bool" },
                        description: { bsonType: "string" }
                      }
                    },
                    diet: {
                      bsonType: "object",
                      properties: {
                        yes: { bsonType: "bool" },
                        description: { bsonType: "string" }
                      }
                    }
                  }
                },
                referralNeeds: {
                  bsonType: "object",
                  properties: {
                    possibleNeeds: { bsonType: "string" },
                    rnAssessment: { bsonType: "string" },
                    date: { bsonType: "string" },
                    time: { bsonType: "string" }
                  }
                }
              }
            },
            createdAt: { bsonType: "date" },
            updatedAt: { bsonType: "date" }
          }
        }
      }
    });

    // hospitals collection
    await db.createCollection('hospitals', {
      validator: {
        $jsonSchema: {
          bsonType: 'object',
          required: ['hospitalId', 'hospitalName', 'hospitalAbbreviation', 'createdAt'],
          properties: {
            hospitalId: { bsonType: 'string' },
            hospitalName: { bsonType: 'string' },
            hospitalAbbreviation: { bsonType: 'string' },
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
    await db.collection('counters').insertOne({ _id: 'patientId', sequence_value: 0 });
    await db.collection('counters').insertOne({ _id: 'hospitalId', sequence_value: 0 });
    await db.collection('counters').insertOne({ _id: 'wardId', sequence_value: 0 });
    await db.collection('counters').insertOne({ _id: 'roomId', sequence_value: 0 });

  },

  async down(db, client) {
    await db.collection('users').drop();
    await db.collection('patients').drop();
    await db.collection('hospitals').drop();
    await db.collection('wards').drop();
    await db.collection('rooms').drop();
    await db.collection('counters').drop();
  }
};
