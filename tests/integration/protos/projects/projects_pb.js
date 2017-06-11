/**
 * @fileoverview
 * @enhanceable
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
goog.exportSymbol('proto.projects.CreateProjectRequest', null, global);
goog.exportSymbol('proto.projects.DeleteProjectRequest', null, global);
goog.exportSymbol('proto.projects.GetProjectRequest', null, global);
goog.exportSymbol('proto.projects.ListProjectsRequest', null, global);
goog.exportSymbol('proto.projects.ListProjectsResponse', null, global);
goog.exportSymbol('proto.projects.Project', null, global);
goog.exportSymbol('proto.projects.UpdateProjectRequest', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.projects.Project = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.projects.Project, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.projects.Project.displayName = 'proto.projects.Project';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.projects.Project.prototype.toObject = function(opt_includeInstance) {
  return proto.projects.Project.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.projects.Project} msg The msg instance to transform.
 * @return {!Object}
 */
proto.projects.Project.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    organizationId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    name: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.projects.Project}
 */
proto.projects.Project.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.projects.Project;
  return proto.projects.Project.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.projects.Project} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.projects.Project}
 */
proto.projects.Project.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setOrganizationId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.projects.Project} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.Project.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.projects.Project.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.Project.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = this.getOrganizationId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = this.getName();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.projects.Project.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.projects.Project.prototype.setId = function(value) {
  jspb.Message.setField(this, 1, value);
};


/**
 * optional string organization_id = 2;
 * @return {string}
 */
proto.projects.Project.prototype.getOrganizationId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.projects.Project.prototype.setOrganizationId = function(value) {
  jspb.Message.setField(this, 2, value);
};


/**
 * optional string name = 3;
 * @return {string}
 */
proto.projects.Project.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.projects.Project.prototype.setName = function(value) {
  jspb.Message.setField(this, 3, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.projects.ListProjectsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.projects.ListProjectsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.projects.ListProjectsRequest.displayName = 'proto.projects.ListProjectsRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.projects.ListProjectsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.projects.ListProjectsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.projects.ListProjectsRequest} msg The msg instance to transform.
 * @return {!Object}
 */
proto.projects.ListProjectsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    pageSize: jspb.Message.getFieldWithDefault(msg, 1, 0),
    pageToken: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.projects.ListProjectsRequest}
 */
proto.projects.ListProjectsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.projects.ListProjectsRequest;
  return proto.projects.ListProjectsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.projects.ListProjectsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.projects.ListProjectsRequest}
 */
proto.projects.ListProjectsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPageSize(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPageToken(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.projects.ListProjectsRequest} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.ListProjectsRequest.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.projects.ListProjectsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.ListProjectsRequest.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getPageSize();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = this.getPageToken();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional int32 page_size = 1;
 * @return {number}
 */
proto.projects.ListProjectsRequest.prototype.getPageSize = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.projects.ListProjectsRequest.prototype.setPageSize = function(value) {
  jspb.Message.setField(this, 1, value);
};


/**
 * optional string page_token = 2;
 * @return {string}
 */
proto.projects.ListProjectsRequest.prototype.getPageToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.projects.ListProjectsRequest.prototype.setPageToken = function(value) {
  jspb.Message.setField(this, 2, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.projects.ListProjectsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.projects.ListProjectsResponse.repeatedFields_, null);
};
goog.inherits(proto.projects.ListProjectsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.projects.ListProjectsResponse.displayName = 'proto.projects.ListProjectsResponse';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.projects.ListProjectsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.projects.ListProjectsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.projects.ListProjectsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.projects.ListProjectsResponse} msg The msg instance to transform.
 * @return {!Object}
 */
proto.projects.ListProjectsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    projectsList: jspb.Message.toObjectList(msg.getProjectsList(),
    proto.projects.Project.toObject, includeInstance),
    nextPageToken: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.projects.ListProjectsResponse}
 */
proto.projects.ListProjectsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.projects.ListProjectsResponse;
  return proto.projects.ListProjectsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.projects.ListProjectsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.projects.ListProjectsResponse}
 */
proto.projects.ListProjectsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.projects.Project;
      reader.readMessage(value,proto.projects.Project.deserializeBinaryFromReader);
      msg.addProjects(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setNextPageToken(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.projects.ListProjectsResponse} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.ListProjectsResponse.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.projects.ListProjectsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.ListProjectsResponse.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getProjectsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.projects.Project.serializeBinaryToWriter
    );
  }
  f = this.getNextPageToken();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * repeated Project projects = 1;
 * If you change this array by adding, removing or replacing elements, or if you
 * replace the array itself, then you must call the setter to update it.
 * @return {!Array.<!proto.projects.Project>}
 */
proto.projects.ListProjectsResponse.prototype.getProjectsList = function() {
  return /** @type{!Array.<!proto.projects.Project>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.projects.Project, 1));
};


/** @param {!Array.<!proto.projects.Project>} value */
proto.projects.ListProjectsResponse.prototype.setProjectsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.projects.Project=} opt_value
 * @param {number=} opt_index
 * @return {!proto.projects.Project}
 */
proto.projects.ListProjectsResponse.prototype.addProjects = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.projects.Project, opt_index);
};


proto.projects.ListProjectsResponse.prototype.clearProjectsList = function() {
  this.setProjectsList([]);
};


/**
 * optional string next_page_token = 2;
 * @return {string}
 */
proto.projects.ListProjectsResponse.prototype.getNextPageToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.projects.ListProjectsResponse.prototype.setNextPageToken = function(value) {
  jspb.Message.setField(this, 2, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.projects.GetProjectRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.projects.GetProjectRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.projects.GetProjectRequest.displayName = 'proto.projects.GetProjectRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.projects.GetProjectRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.projects.GetProjectRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.projects.GetProjectRequest} msg The msg instance to transform.
 * @return {!Object}
 */
proto.projects.GetProjectRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.projects.GetProjectRequest}
 */
proto.projects.GetProjectRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.projects.GetProjectRequest;
  return proto.projects.GetProjectRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.projects.GetProjectRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.projects.GetProjectRequest}
 */
proto.projects.GetProjectRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.projects.GetProjectRequest} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.GetProjectRequest.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.projects.GetProjectRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.GetProjectRequest.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.projects.GetProjectRequest.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.projects.GetProjectRequest.prototype.setId = function(value) {
  jspb.Message.setField(this, 1, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.projects.CreateProjectRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.projects.CreateProjectRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.projects.CreateProjectRequest.displayName = 'proto.projects.CreateProjectRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.projects.CreateProjectRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.projects.CreateProjectRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.projects.CreateProjectRequest} msg The msg instance to transform.
 * @return {!Object}
 */
proto.projects.CreateProjectRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    project: (f = msg.getProject()) && proto.projects.Project.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.projects.CreateProjectRequest}
 */
proto.projects.CreateProjectRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.projects.CreateProjectRequest;
  return proto.projects.CreateProjectRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.projects.CreateProjectRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.projects.CreateProjectRequest}
 */
proto.projects.CreateProjectRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.projects.Project;
      reader.readMessage(value,proto.projects.Project.deserializeBinaryFromReader);
      msg.setProject(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.projects.CreateProjectRequest} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.CreateProjectRequest.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.projects.CreateProjectRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.CreateProjectRequest.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getProject();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.projects.Project.serializeBinaryToWriter
    );
  }
};


/**
 * optional Project project = 1;
 * @return {?proto.projects.Project}
 */
proto.projects.CreateProjectRequest.prototype.getProject = function() {
  return /** @type{?proto.projects.Project} */ (
    jspb.Message.getWrapperField(this, proto.projects.Project, 1));
};


/** @param {?proto.projects.Project|undefined} value */
proto.projects.CreateProjectRequest.prototype.setProject = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.projects.CreateProjectRequest.prototype.clearProject = function() {
  this.setProject(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.projects.CreateProjectRequest.prototype.hasProject = function() {
  return jspb.Message.getField(this, 1) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.projects.UpdateProjectRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.projects.UpdateProjectRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.projects.UpdateProjectRequest.displayName = 'proto.projects.UpdateProjectRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.projects.UpdateProjectRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.projects.UpdateProjectRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.projects.UpdateProjectRequest} msg The msg instance to transform.
 * @return {!Object}
 */
proto.projects.UpdateProjectRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    project: (f = msg.getProject()) && proto.projects.Project.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.projects.UpdateProjectRequest}
 */
proto.projects.UpdateProjectRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.projects.UpdateProjectRequest;
  return proto.projects.UpdateProjectRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.projects.UpdateProjectRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.projects.UpdateProjectRequest}
 */
proto.projects.UpdateProjectRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = new proto.projects.Project;
      reader.readMessage(value,proto.projects.Project.deserializeBinaryFromReader);
      msg.setProject(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.projects.UpdateProjectRequest} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.UpdateProjectRequest.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.projects.UpdateProjectRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.UpdateProjectRequest.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = this.getProject();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.projects.Project.serializeBinaryToWriter
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.projects.UpdateProjectRequest.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.projects.UpdateProjectRequest.prototype.setId = function(value) {
  jspb.Message.setField(this, 1, value);
};


/**
 * optional Project project = 2;
 * @return {?proto.projects.Project}
 */
proto.projects.UpdateProjectRequest.prototype.getProject = function() {
  return /** @type{?proto.projects.Project} */ (
    jspb.Message.getWrapperField(this, proto.projects.Project, 2));
};


/** @param {?proto.projects.Project|undefined} value */
proto.projects.UpdateProjectRequest.prototype.setProject = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.projects.UpdateProjectRequest.prototype.clearProject = function() {
  this.setProject(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.projects.UpdateProjectRequest.prototype.hasProject = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.projects.DeleteProjectRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.projects.DeleteProjectRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.projects.DeleteProjectRequest.displayName = 'proto.projects.DeleteProjectRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.projects.DeleteProjectRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.projects.DeleteProjectRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.projects.DeleteProjectRequest} msg The msg instance to transform.
 * @return {!Object}
 */
proto.projects.DeleteProjectRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.projects.DeleteProjectRequest}
 */
proto.projects.DeleteProjectRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.projects.DeleteProjectRequest;
  return proto.projects.DeleteProjectRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.projects.DeleteProjectRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.projects.DeleteProjectRequest}
 */
proto.projects.DeleteProjectRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.projects.DeleteProjectRequest} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.DeleteProjectRequest.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.projects.DeleteProjectRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.projects.DeleteProjectRequest.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.projects.DeleteProjectRequest.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.projects.DeleteProjectRequest.prototype.setId = function(value) {
  jspb.Message.setField(this, 1, value);
};


goog.object.extend(exports, proto.projects);
