package main

var schemas = `
{
    "API": {
        "createAsset": {
            "description": "Create an asset. One argument, a JSON encoded event. farmID is required with zero or more writable properties. Establishes an initial asset state.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "A set of fields that constitute the writable fields in an asset's state. farmID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                        "properties": {
                            "farmID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "farmerid": {
                                "description": "transport entity currently in possession of asset",
                                "type": "string"
                            },
                            "location": {
                                "description": "A geographical coordinate",
                                "properties": {
                                    "latitude": {
                                        "type": "number"
                                    },
                                    "longitude": {
                                        "type": "number"
                                    }
                                },
                                "type": "object"
                            },
							 "watersalinity": {
                                "description": "person watersalinity",
                                "type": "number"
                            },
							"soilph": {
                                "description": "soilph name",
                                "type": "number"
                            },
							 "soilmoisture": {
                                "description": "overall soilmoisture",
                                "type": "number"
                            },
							 "nitrogen": {
                                "description": "nitrogen",
                                "type": "number"
                            },
							 "potassium": {
                                "description": "potassium",
                                "type": "number"
                            },
							 "sulphur": {
                                "description": "sulphur",
                                "type": "number"
                            },
							"ambienttemperature": {
                                "description": "ambienttemperature",
                                "type": "number"
                            },
							"ambienthumidity": {
                                "description": "ambienthumidity",
                                "type": "number"
                            },
							"soiltemperature": {
                                "description": "soiltemperature",
                                "type": "number"
                            },
							"luminosity": {
                                "description": "luminosity",
                                "type": "number"
                            },
							"time": {
                                "description": "time",
                                "type": "number"
                            },
							"soilindex": {
                                "description": "soilindex",
                                "type": "number"
                            }
                        },
                        "required": [
                            "farmID"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "createAsset function",
                    "enum": [
                        "createAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAsset": {
            "description": "Delete an asset. Argument is a JSON encoded string containing only an farmID.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an farmID for use as an argument to read or delete.",
                        "properties": {
                            "farmID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "deleteAsset function",
                    "enum": [
                        "deleteAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "init": {
            "description": "Initializes the contract when started, either by deployment or by peer restart.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "event sent to init on deployment",
                        "properties": {
                            "nickname": {
                                "default": "SIMPLE",
                                "description": "The nickname of the current contract",
                                "type": "string"
                            },
                            "version": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "required": [
                            "version"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "init function",
                    "enum": [
                        "init"
                    ],
                    "type": "string"
                },
                "method": "deploy"
            },
            "type": "object"
        },
        "readAsset": {
            "description": "Returns the state an asset. Argument is a JSON encoded string. farmID is the only accepted property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an farmID for use as an argument to read or delete.",
                        "properties": {
                            "farmID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "readAsset function",
                    "enum": [
                        "readAsset"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of fields that constitute the complete asset state.",
                    "properties": {
                        "farmID": {
                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                            "type": "string"
                        },
                        "farmerid": {
                            "description": "transport entity currently in possession of asset",
                            "type": "string"
                        },
                        "location": {
                            "description": "A geographical coordinate",
                            "properties": {
                                "latitude": {
                                    "type": "number"
                                },
                                "longitude": {
                                    "type": "number"
                                }
                            },
                            "type": "object"
                        },
							 "watersalinity": {
                                "description": "person watersalinity",
                                "type": "number"
                            },
							"soilph": {
                                "description": "soilph name",
                                "type": "number"
                            },
							 "soilmoisture": {
                                "description": "overall soilmoisture",
                                "type": "number"
                            },
							 "nitrogen": {
                                "description": "nitrogen",
                                "type": "number"
                            },
							 "potassium": {
                                "description": "potassium",
                                "type": "number"
                            },
							 "sulphur": {
                                "description": "sulphur",
                                "type": "number"
                            },
							"ambienttemperature": {
                                "description": "ambienttemperature",
                                "type": "number"
                            },
							"ambienthumidity": {
                                "description": "ambienthumidity",
                                "type": "number"
                            },
							"soiltemperature": {
                                "description": "soiltemperature",
                                "type": "number"
                            },
							"luminosity": {
                                "description": "luminosity",
                                "type": "number"
                            },
							"time": {
                                "description": "time",
                                "type": "number"
                            },
							"soilindex": {
                                "description": "soilindex",
                                "type": "number"
                            }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetHistory": {
            "description": "Requests a specified number of history states for an assets. Returns an array of states sorted with the most recent first. farmID is required and count is optional. A missing count, a count of zero, or too large a count returns all existing history states.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested farmID",
                        "properties": {
                            "farmID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "required": [
                            "farmID"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "readAssetHistory function",
                    "enum": [
                        "readAssetHistory"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "an array of states for one asset sorted by timestamp with the most recent entry first",
                    "items": {
                        "description": "A set of fields that constitute the complete asset state.",
                        "properties": {
                            "farmID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "farmerid": {
                                "description": "transport entity currently in possession of asset",
                                "type": "string"
                            },
                            "location": {
                                "description": "A geographical coordinate",
                                "properties": {
                                    "latitude": {
                                        "type": "number"
                                    },
                                    "longitude": {
                                        "type": "number"
                                    }
                                },
                                "type": "object"
                            },
							 "watersalinity": {
                                "description": "person watersalinity",
                                "type": "number"
                            },
							"soilph": {
                                "description": "soilph name",
                                "type": "number"
                            },
							 "soilmoisture": {
                                "description": "overall soilmoisture",
                                "type": "number"
                            },
							 "nitrogen": {
                                "description": "nitrogen",
                                "type": "number"
                            },
							 "potassium": {
                                "description": "potassium",
                                "type": "number"
                            },
							 "sulphur": {
                                "description": "sulphur",
                                "type": "number"
                            },
							"ambienttemperature": {
                                "description": "ambienttemperature",
                                "type": "number"
                            },
							"ambienthumidity": {
                                "description": "ambienthumidity",
                                "type": "number"
                            },
							"soiltemperature": {
                                "description": "soiltemperature",
                                "type": "number"
                            },
							"luminosity": {
                                "description": "luminosity",
                                "type": "number"
                            },
							"time": {
                                "description": "time",
                                "type": "number"
                            },
							"soilindex": {
                                "description": "soilindex",
                                "type": "number"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAssetSamples": {
            "description": "Returns a string generated from the schema containing sample Objects as specified in generate.json in the scripts folder.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readAssetSamples function",
                    "enum": [
                        "readAssetSamples"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "JSON encoded object containing selected sample data",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "readAssetSchemas": {
            "description": "Returns a string generated from the schema containing APIs and Objects as specified in generate.json in the scripts folder.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readAssetSchemas function",
                    "enum": [
                        "readAssetSchemas"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "JSON encoded object containing selected schemas",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "updateAsset": {
            "description": "Update the state of an asset. The one argument is a JSON encoded event. farmID is required along with one or more writable properties. Establishes the next asset state. ",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "A set of fields that constitute the writable fields in an asset's state. farmID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                        "properties": {
                            "farmID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "farmerid": {
                                "description": "transport entity currently in possession of asset",
                                "type": "string"
                            },
                            "location": {
                                "description": "A geographical coordinate",
                                "properties": {
                                    "latitude": {
                                        "type": "number"
                                    },
                                    "longitude": {
                                        "type": "number"
                                    }
                                },
                                "type": "object"
                            },
							 "watersalinity": {
                                "description": "person watersalinity",
                                "type": "number"
                            },
							"soilph": {
                                "description": "soilph name",
                                "type": "number"
                            },
							 "soilmoisture": {
                                "description": "overall soilmoisture",
                                "type": "number"
                            },
							 "nitrogen": {
                                "description": "nitrogen",
                                "type": "number"
                            },
							 "potassium": {
                                "description": "potassium",
                                "type": "number"
                            },
							 "sulphur": {
                                "description": "sulphur",
                                "type": "number"
                            },
							"ambienttemperature": {
                                "description": "ambienttemperature",
                                "type": "number"
                            },
							"ambienthumidity": {
                                "description": "ambienthumidity",
                                "type": "number"
                            },
							"soiltemperature": {
                                "description": "soiltemperature",
                                "type": "number"
                            },
							"luminosity": {
                                "description": "luminosity",
                                "type": "number"
                            },
							"time": {
                                "description": "time",
                                "type": "number"
                            },
							"soilindex": {
                                "description": "soilindex",
                                "type": "number"
                            }
                        },
                        "required": [
                            "farmID"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "updateAsset function",
                    "enum": [
                        "updateAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        }
    },
    "objectModelSchemas": {
        "farmIDKey": {
            "description": "An object containing only an farmID for use as an argument to read or delete.",
            "properties": {
                "farmID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "event": {
            "description": "A set of fields that constitute the writable fields in an asset's state. farmID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
            "properties": {
                "farmID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "farmerid": {
                    "description": "transport entity currently in possession of asset",
                    "type": "string"
                },
                "location": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "number"
                        },
                        "longitude": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                },"watersalinity": {
                                "description": "person watersalinity",
                                "type": "number"
                            },
							"soilph": {
                                "description": "soilph name",
                                "type": "number"
                            },
							 "soilmoisture": {
                                "description": "overall soilmoisture",
                                "type": "number"
                            },
							 "nitrogen": {
                                "description": "nitrogen",
                                "type": "number"
                            },
							 "potassium": {
                                "description": "potassium",
                                "type": "number"
                            },
							 "sulphur": {
                                "description": "sulphur",
                                "type": "number"
                            },
							"ambienttemperature": {
                                "description": "ambienttemperature",
                                "type": "number"
                            },
							"ambienthumidity": {
                                "description": "ambienthumidity",
                                "type": "number"
                            },
							"soiltemperature": {
                                "description": "soiltemperature",
                                "type": "number"
                            },
							"luminosity": {
                                "description": "luminosity",
                                "type": "number"
                            },
							"time": {
                                "description": "time",
                                "type": "number"
                            },
							"soilindex": {
                                "description": "soilindex",
                                "type": "number"
                            }
            },
            "required": [
                "farmID"
            ],
            "type": "object"
        },
        "initEvent": {
            "description": "event sent to init on deployment",
            "properties": {
                "nickname": {
                    "default": "SIMPLE",
                    "description": "The nickname of the current contract",
                    "type": "string"
                },
                "version": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "required": [
                "version"
            ],
            "type": "object"
        },
        "state": {
            "description": "A set of fields that constitute the complete asset state.",
            "properties": {
                "farmID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "farmerid": {
                    "description": "transport entity currently in possession of asset",
                    "type": "string"
                },
                "location": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "string"
                        },
                        "longitude": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
							 "watersalinity": {
                                "description": "person watersalinity",
                                "type": "number"
                            },
							"soilph": {
                                "description": "soilph name",
                                "type": "number"
                            },
							 "soilmoisture": {
                                "description": "overall soilmoisture",
                                "type": "number"
                            },
							 "nitrogen": {
                                "description": "nitrogen",
                                "type": "number"
                            },
							 "potassium": {
                                "description": "potassium",
                                "type": "number"
                            },
							 "sulphur": {
                                "description": "sulphur",
                                "type": "number"
                            },
							"ambienttemperature": {
                                "description": "ambienttemperature",
                                "type": "number"
                            },"ambienthumidity": {
                                "description": "ambienthumidity",
                                "type": "number"
                            },"soiltemperature": {
                                "description": "soiltemperature",
                                "type": "number"
                            },"luminosity": {
                                "description": "luminosity",
                                "type": "number"
                            },
							"time": {
                                "description": "time",
                                "type": "number"
                            },
							"soilindex": {
                                "description": "soilindex",
                                "type": "number"
                            }
            },
            "type": "object"
        }
    }
}`
