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
            dateOfBirth: { bsonType: 'string' },
            height: { bsonType: 'double' },
            weight: { bsonType: 'double' },
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

    // chart_reviews collection
    await db.createCollection('chart_reviews', {
      validator: {
        $jsonSchema: {
          bsonType: 'object',
          required: ['patientId', 'createdAt', 'updatedAt'],
          properties: {
            patientId: { bsonType: 'string' },
            charts: {
              bsonType: 'array',
              items: {
                bsonType: 'object',
                required: ['chartType', 'top', 'left'],
                properties: {
                  chartType: { bsonType: 'string' },
                  top: { bsonType: 'string' },
                  left: { bsonType: 'string' },
                },
              },
            },
            createdAt: { bsonType: 'date' },
            updatedAt: { bsonType: 'date' },
          },
        },
      },
    });

    // patient_assessments collection
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
            skin: {
              bsonType: "object",
              properties: {
                front: {
                  bsonType: "array",
                  items: {
                    bsonType: "object",
                    properties: {
                      skinType: { bsonType: "string" },
                      description: { bsonType: "string" },
                      top: { bsonType: "string" },
                      left: { bsonType: "string" }
                    }
                  }
                },
                back: {
                  bsonType: "array",
                  items: {
                    bsonType: "object",
                    properties: {
                      skinType: { bsonType: "string" },
                      description: { bsonType: "string" },
                      top: { bsonType: "string" },
                      left: { bsonType: "string" }
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

    // nurse_monitorings collection
    await db.createCollection('nurse_monitorings', {
      validator: {
        $jsonSchema: {
          bsonType: "object",
          required: ["patientId", "createdAt", "updatedAt"],
          properties: {
            patientId: { bsonType: "string" },
            heartRate: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  value: { bsonType: "int" },
                  recordedAt: { bsonType: "date" }
                }
              }
            },
            fetalHeartRate: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  value: { bsonType: "int" },
                  recordedAt: { bsonType: "date" }
                }
              }
            },
            bloodPressure: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  systol: { bsonType: "int" },
                  diastol: { bsonType: "int" },
                  recordedAt: { bsonType: "date" }
                }
              }
            },
            respiration: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  value: { bsonType: "int" },
                  recordedAt: { bsonType: "date" }
                }
              }
            },
            oxygenSaturation: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  value: { bsonType: "int" },
                  recordedAt: { bsonType: "date" }
                }
              }
            },
            temperature: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  value: { bsonType: "double" },
                  recordedAt: { bsonType: "date" }
                }
              }
            },
            painScale: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  value: { bsonType: "int" },
                  recordedAt: { bsonType: "date" }
                }
              }
            },
            neurological: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  levelOfConsciousness: { bsonType: "string" },
                  sendationScore: { bsonType: "int" },
                  pupilSize: { bsonType: "int" },
                  pupilReaction: { bsonType: "string" },
                  recordedAt: { bsonType: "date" }
                }
              }
            },
            fluidIntakeOutput: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  shiftID: { bsonType: "string" },
                  startDateTime: { bsonType: "date" },
                  endDateTime: { bsonType: ["date", "null"] },
                  method: { bsonType: "string" },
                  site: { bsonType: "string" },
                  shift: {
                    bsonType: "object",
                    properties: {
                      rows: {
                        bsonType: "array",
                        items: {
                          bsonType: "object",
                          properties: {
                            time: { bsonType: "string" },
                            typeOfFluid: { bsonType: "string" },
                            additionsPerBag: { bsonType: "string" },
                            putUp: { bsonType: "int" },
                            goneIn: { bsonType: "int" },
                            urine: { bsonType: "int" },
                            ngAspirate: { bsonType: "int" },
                            drains: { bsonType: "int" },
                            stool: { bsonType: "string" }
                          }
                        }
                      },
                      endOfShiftTime: { bsonType: "string" },
                      remainder: { bsonType: "int" }
                    }
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

    // medications collection
    await db.createCollection('medications', {
      validator: {
        $jsonSchema: {
          bsonType: "object",
          required: ["patientId", "createdAt", "updatedAt"],
          properties: {
            patientId: { bsonType: "string" },
            injection: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  medication: { bsonType: "string" },
                  dose: { bsonType: "string" },
                  route: { bsonType: "string" },
                  frequency: { bsonType: "date" }
                }
              }
            },
            intravenous: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  medication: { bsonType: "string" },
                  dose: { bsonType: "string" },
                  route: { bsonType: "string" },
                  frequency: { bsonType: "date" }
                }
              }
            },
            oral: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  medication: { bsonType: "string" },
                  dose: { bsonType: "string" },
                  frequency: { bsonType: "date" }
                }
              }
            },
            tropical: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  medication: { bsonType: "string" },
                  dose: { bsonType: "string" },
                  route: { bsonType: "string" },
                  frequency: { bsonType: "date" }
                }
              }
            },
            drop: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  medication: { bsonType: "string" },
                  dose: { bsonType: "string" },
                  route: { bsonType: "string" },
                  frequency: { bsonType: "date" }
                }
              }
            },
            implant: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  medication: { bsonType: "string" },
                  dose: { bsonType: "string" },
                  route: { bsonType: "string" },
                  frequency: { bsonType: "date" }
                }
              }
            },
            suppositories: {
              bsonType: "array",
              items: {
                bsonType: "object",
                properties: {
                  medication: { bsonType: "string" },
                  dose: { bsonType: "string" },
                  frequency: { bsonType: "date" }
                }
              }
            },
            createdAt: { bsonType: "date" },
            updatedAt: { bsonType: "date" }
          }
        }
      }
    });

    // nursing collection
    await db.createCollection('nursings', {
      validator: {
        $jsonSchema: {
          bsonType: "object",
          required: [
            "patientId",
            "activityFlowSheet",
            "focusNote",
            "focusList",
            "dischargeSummaryForm",
            "createdAt",
            "updatedAt"
          ],
          properties: {
            patientId: { bsonType: "string" },
            activityFlowSheet: {
              bsonType: "object",
              required: [
                "specialCare",
                "hygiene",
                "activity",
                "psychosocialSupport",
                "elimination",
                "safety"
              ],
              properties: {
                specialCare: {
                  bsonType: "object",
                  required: [
                    "ivfIVCathCare",
                    "drainCare",
                    "woundCare",
                    "painCare",
                    "breathingExercise",
                    "turnPosition",
                    "other"
                  ],
                  properties: {
                    ivfIVCathCare: { bsonType: "string" },
                    drainCare: { bsonType: "string" },
                    woundCare: { bsonType: "string" },
                    painCare: { bsonType: "string" },
                    breathingExercise: { bsonType: "string" },
                    turnPosition: { bsonType: "string" },
                    other: { bsonType: "string" }
                  }
                },
                hygiene: {
                  bsonType: "object",
                  required: [
                    "bath",
                    "mouthCare",
                    "shampooing",
                    "perinealCare",
                    "linenChange",
                    "other"
                  ],
                  properties: {
                    bath: { bsonType: "string" },
                    mouthCare: { bsonType: "string" },
                    shampooing: { bsonType: "string" },
                    perinealCare: { bsonType: "string" },
                    linenChange: { bsonType: "string" },
                    other: { bsonType: "string" }
                  }
                },
                activity: {
                  bsonType: "object",
                  required: [
                    "ambulate",
                    "exercise",
                    "rest",
                    "bedRest",
                    "absoluteBedRest"
                  ],
                  properties: {
                    ambulate: { bsonType: "string" },
                    exercise: { bsonType: "string" },
                    rest: { bsonType: "string" },
                    bedRest: { bsonType: "string" },
                    absoluteBedRest: { bsonType: "string" }
                  }
                },
                psychosocialSupport: {
                  bsonType: "object",
                  required: [
                    "greeting",
                    "touch:",
                    "information",
                    "other"
                  ],
                  properties: {
                    greeting: { bsonType: "string" },
                    // Note: The property name "touch:" includes a colon.
                    "touch:": { bsonType: "string" },
                    information: { bsonType: "string" },
                    other: { bsonType: "string" }
                  }
                },
                elimination: {
                  bsonType: "object",
                  required: [
                    "condomCare",
                    "cathCare",
                    "ostomyCare",
                    "enema"
                  ],
                  properties: {
                    condomCare: { bsonType: "string" },
                    cathCare: { bsonType: "string" },
                    ostomyCare: { bsonType: "string" },
                    enema: { bsonType: "string" }
                  }
                },
                safety: {
                  bsonType: "object",
                  required: [
                    "sideRailUp",
                    "restraint",
                    "other"
                  ],
                  properties: {
                    sideRailUp: { bsonType: "string" },
                    restraint: { bsonType: "string" },
                    other: { bsonType: "string" }
                  }
                }
              }
            },
            focusNote: {
              bsonType: "array",
              items: {
                bsonType: "object",
                required: [
                  "timeOrShift",
                  "focus",
                  "progressNote",
                  "createdAt",
                  "updatedAt"
                ],
                properties: {
                  timeOrShift: { bsonType: "string" },
                  focus: { bsonType: "string" },
                  progressNote: { bsonType: "string" },
                  createdAt: { bsonType: "date" },
                  updatedAt: { bsonType: "date" }
                }
              }
            },
            focusList: {
              bsonType: "array",
              items: {
                bsonType: "object",
                required: [
                  "no",
                  "focus",
                  "goal",
                  "active",
                  "resolved"
                ],
                properties: {
                  no: { bsonType: "string" },
                  focus: { bsonType: "string" },
                  goal: { bsonType: "string" },
                  active: { bsonType: "string" },
                  resolved: { bsonType: "string" }
                }
              }
            },
            dischargeSummaryForm: {
              bsonType: "object",
              required: [
                "doctor",
                "diagnosis",
                "treatment",
                "status",
                "continuingHealthProblems",
                "healthInformations",
                "dischargedBy",
                "dischargedWith",
                "caseOfDeath"
              ],
              properties: {
                doctor: { bsonType: "string" },
                diagnosis: { bsonType: "string" },
                treatment: { bsonType: "string" },
                status: {
                  bsonType: "object",
                  required: [
                    "vitalSigns",
                    "levelOfConciousness",
                    "therapeuticDevice"
                  ],
                  properties: {
                    vitalSigns: {
                      bsonType: "object",
                      required: [
                        "temperature",
                        "heartRate",
                        "respiratoryRate",
                        "bloodPressure"
                      ],
                      properties: {
                        temperature: { bsonType: "double" },
                        heartRate: { bsonType: "int" },
                        respiratoryRate: { bsonType: "int" },
                        bloodPressure: {
                          bsonType: "object",
                          required: ["systol", "diastol"],
                          properties: {
                            systol: { bsonType: "int" },
                            diastol: { bsonType: "int" }
                          }
                        }
                      }
                    },
                    levelOfConciousness: { bsonType: "string" },
                    therapeuticDevice: { bsonType: "string" }
                  }
                },
                continuingHealthProblems: { bsonType: "string" },
                healthInformations: {
                  bsonType: "object",
                  required: [
                    "selected",
                    "additionalInfo",
                    "providedTo"
                  ],
                  properties: {
                    selected: { bsonType: "string" },
                    additionalInfo: { bsonType: "string" },
                    providedTo: { bsonType: "string" }
                  }
                },
                dischargedBy: { bsonType: "string" },
                dischargedWith: { bsonType: "string" },
                caseOfDeath: {
                  bsonType: "object",
                  required: [
                    "dateOfDeath",
                    "timeOfDeath",
                    "doctor",
                    "transferTo",
                    "propertySentTo",
                    "dischargeRN"
                  ],
                  properties: {
                    dateOfDeath: { bsonType: "string" },
                    timeOfDeath: { bsonType: "string" },
                    doctor: { bsonType: "string" },
                    transferTo: { bsonType: "string" },
                    propertySentTo: {
                      bsonType: "object",
                      required: [
                        "selected",
                        "signature",
                        "idNo"
                      ],
                      properties: {
                        selected: { bsonType: "string" },
                        signature: { bsonType: "string" },
                        idNo: { bsonType: "string" }
                      }
                    },
                    dischargeRN: { bsonType: "string" }
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
    await db.collection('chart_reviews')
    await db.collection('patient_assessments').drop();
    await db.collection('nurse_monitorings').drop();
    await db.collection('medications').drop();
    await db.collection('nursings').drop();
    await db.collection('hospitals').drop();
    await db.collection('wards').drop();
    await db.collection('rooms').drop();
    await db.collection('counters').drop();
  }
};
