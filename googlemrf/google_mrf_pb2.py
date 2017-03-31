# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: google_mrf.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='google_mrf.proto',
  package='googlemrf',
  syntax='proto3',
  serialized_pb=_b('\n\x10google_mrf.proto\x12\tgooglemrf\"\x07\n\x05\x45mpty\"%\n\x05Query\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06parent\x18\x02 \x01(\t\"\xc8\x01\n\x03MRF\x12\x11\n\tSubmitter\x18\x01 \x01(\t\x12\n\n\x02ID\x18\x02 \x01(\t\x12\x0c\n\x04Name\x18\x03 \x01(\t\x12\x0f\n\x07Product\x18\x04 \x01(\t\x12\x13\n\x0bProductCode\x18\x05 \x01(\t\x12\x12\n\nSubProduct\x18\x06 \x01(\t\x12\x12\n\nCostCenter\x18\x07 \x01(\t\x12\x15\n\rPrimaryRegion\x18\x08 \x01(\t\x12\x0c\n\x04Year\x18\t \x01(\x05\x12\x0f\n\x07Quarter\x18\n \x01(\x05\x12\x10\n\x08\x41pproved\x18\x0b \x01(\x08\"\x16\n\x06Single\x12\x0c\n\x04Name\x18\x01 \x01(\t\"&\n\x06\x44ouble\x12\x0c\n\x04Name\x18\x01 \x01(\t\x12\x0e\n\x06Parent\x18\x02 \x01(\t2\x9f\x07\n\tGoogleMRF\x12-\n\x07MRFInfo\x12\x10.googlemrf.Query\x1a\x0e.googlemrf.MRF\"\x00\x12,\n\x04MRFs\x12\x10.googlemrf.Empty\x1a\x0e.googlemrf.MRF\"\x00\x30\x01\x12\x33\n\x08\x41gencies\x12\x10.googlemrf.Empty\x1a\x11.googlemrf.Single\"\x00\x30\x01\x12/\n\x04LOBs\x12\x10.googlemrf.Empty\x1a\x11.googlemrf.Single\"\x00\x30\x01\x12\x33\n\x08Products\x12\x10.googlemrf.Empty\x1a\x11.googlemrf.Double\"\x00\x30\x01\x12\x36\n\x0bSubProducts\x12\x10.googlemrf.Empty\x1a\x11.googlemrf.Double\"\x00\x30\x01\x12\x33\n\x08\x43hannels\x12\x10.googlemrf.Empty\x1a\x11.googlemrf.Single\"\x00\x30\x01\x12\x31\n\x06Medias\x12\x10.googlemrf.Empty\x1a\x11.googlemrf.Double\"\x00\x30\x01\x12\x34\n\tSubMedias\x12\x10.googlemrf.Empty\x1a\x11.googlemrf.Double\"\x00\x30\x01\x12\x31\n\x0bValidateMRF\x12\x10.googlemrf.Query\x1a\x0e.googlemrf.MRF\"\x00\x12\x37\n\x0eValidateAgency\x12\x10.googlemrf.Query\x1a\x11.googlemrf.Single\"\x00\x12\x38\n\x0fValidateChannel\x12\x10.googlemrf.Query\x1a\x11.googlemrf.Single\"\x00\x12\x34\n\x0bValidateLOB\x12\x10.googlemrf.Query\x1a\x11.googlemrf.Single\"\x00\x12\x36\n\rValidateMedia\x12\x10.googlemrf.Query\x1a\x11.googlemrf.Single\"\x00\x12\x38\n\x0fValidateProduct\x12\x10.googlemrf.Query\x1a\x11.googlemrf.Single\"\x00\x12\x39\n\x10ValidateSubMedia\x12\x10.googlemrf.Query\x1a\x11.googlemrf.Single\"\x00\x12;\n\x12ValidateSubProduct\x12\x10.googlemrf.Query\x1a\x11.googlemrf.Single\"\x00\x62\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_EMPTY = _descriptor.Descriptor(
  name='Empty',
  full_name='googlemrf.Empty',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=31,
  serialized_end=38,
)


_QUERY = _descriptor.Descriptor(
  name='Query',
  full_name='googlemrf.Query',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='googlemrf.Query.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='parent', full_name='googlemrf.Query.parent', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=40,
  serialized_end=77,
)


_MRF = _descriptor.Descriptor(
  name='MRF',
  full_name='googlemrf.MRF',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Submitter', full_name='googlemrf.MRF.Submitter', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ID', full_name='googlemrf.MRF.ID', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Name', full_name='googlemrf.MRF.Name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Product', full_name='googlemrf.MRF.Product', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ProductCode', full_name='googlemrf.MRF.ProductCode', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='SubProduct', full_name='googlemrf.MRF.SubProduct', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='CostCenter', full_name='googlemrf.MRF.CostCenter', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='PrimaryRegion', full_name='googlemrf.MRF.PrimaryRegion', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Year', full_name='googlemrf.MRF.Year', index=8,
      number=9, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Quarter', full_name='googlemrf.MRF.Quarter', index=9,
      number=10, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Approved', full_name='googlemrf.MRF.Approved', index=10,
      number=11, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=80,
  serialized_end=280,
)


_SINGLE = _descriptor.Descriptor(
  name='Single',
  full_name='googlemrf.Single',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Name', full_name='googlemrf.Single.Name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=282,
  serialized_end=304,
)


_DOUBLE = _descriptor.Descriptor(
  name='Double',
  full_name='googlemrf.Double',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Name', full_name='googlemrf.Double.Name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Parent', full_name='googlemrf.Double.Parent', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=306,
  serialized_end=344,
)

DESCRIPTOR.message_types_by_name['Empty'] = _EMPTY
DESCRIPTOR.message_types_by_name['Query'] = _QUERY
DESCRIPTOR.message_types_by_name['MRF'] = _MRF
DESCRIPTOR.message_types_by_name['Single'] = _SINGLE
DESCRIPTOR.message_types_by_name['Double'] = _DOUBLE

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), dict(
  DESCRIPTOR = _EMPTY,
  __module__ = 'google_mrf_pb2'
  # @@protoc_insertion_point(class_scope:googlemrf.Empty)
  ))
_sym_db.RegisterMessage(Empty)

Query = _reflection.GeneratedProtocolMessageType('Query', (_message.Message,), dict(
  DESCRIPTOR = _QUERY,
  __module__ = 'google_mrf_pb2'
  # @@protoc_insertion_point(class_scope:googlemrf.Query)
  ))
_sym_db.RegisterMessage(Query)

MRF = _reflection.GeneratedProtocolMessageType('MRF', (_message.Message,), dict(
  DESCRIPTOR = _MRF,
  __module__ = 'google_mrf_pb2'
  # @@protoc_insertion_point(class_scope:googlemrf.MRF)
  ))
_sym_db.RegisterMessage(MRF)

Single = _reflection.GeneratedProtocolMessageType('Single', (_message.Message,), dict(
  DESCRIPTOR = _SINGLE,
  __module__ = 'google_mrf_pb2'
  # @@protoc_insertion_point(class_scope:googlemrf.Single)
  ))
_sym_db.RegisterMessage(Single)

Double = _reflection.GeneratedProtocolMessageType('Double', (_message.Message,), dict(
  DESCRIPTOR = _DOUBLE,
  __module__ = 'google_mrf_pb2'
  # @@protoc_insertion_point(class_scope:googlemrf.Double)
  ))
_sym_db.RegisterMessage(Double)


try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces


  class GoogleMRFStub(object):
    """Interface exported by the server.
    """

    def __init__(self, channel):
      """Constructor.

      Args:
        channel: A grpc.Channel.
      """
      self.MRFInfo = channel.unary_unary(
          '/googlemrf.GoogleMRF/MRFInfo',
          request_serializer=Query.SerializeToString,
          response_deserializer=MRF.FromString,
          )
      self.MRFs = channel.unary_stream(
          '/googlemrf.GoogleMRF/MRFs',
          request_serializer=Empty.SerializeToString,
          response_deserializer=MRF.FromString,
          )
      self.Agencies = channel.unary_stream(
          '/googlemrf.GoogleMRF/Agencies',
          request_serializer=Empty.SerializeToString,
          response_deserializer=Single.FromString,
          )
      self.LOBs = channel.unary_stream(
          '/googlemrf.GoogleMRF/LOBs',
          request_serializer=Empty.SerializeToString,
          response_deserializer=Single.FromString,
          )
      self.Products = channel.unary_stream(
          '/googlemrf.GoogleMRF/Products',
          request_serializer=Empty.SerializeToString,
          response_deserializer=Double.FromString,
          )
      self.SubProducts = channel.unary_stream(
          '/googlemrf.GoogleMRF/SubProducts',
          request_serializer=Empty.SerializeToString,
          response_deserializer=Double.FromString,
          )
      self.Channels = channel.unary_stream(
          '/googlemrf.GoogleMRF/Channels',
          request_serializer=Empty.SerializeToString,
          response_deserializer=Single.FromString,
          )
      self.Medias = channel.unary_stream(
          '/googlemrf.GoogleMRF/Medias',
          request_serializer=Empty.SerializeToString,
          response_deserializer=Double.FromString,
          )
      self.SubMedias = channel.unary_stream(
          '/googlemrf.GoogleMRF/SubMedias',
          request_serializer=Empty.SerializeToString,
          response_deserializer=Double.FromString,
          )
      self.ValidateMRF = channel.unary_unary(
          '/googlemrf.GoogleMRF/ValidateMRF',
          request_serializer=Query.SerializeToString,
          response_deserializer=MRF.FromString,
          )
      self.ValidateAgency = channel.unary_unary(
          '/googlemrf.GoogleMRF/ValidateAgency',
          request_serializer=Query.SerializeToString,
          response_deserializer=Single.FromString,
          )
      self.ValidateChannel = channel.unary_unary(
          '/googlemrf.GoogleMRF/ValidateChannel',
          request_serializer=Query.SerializeToString,
          response_deserializer=Single.FromString,
          )
      self.ValidateLOB = channel.unary_unary(
          '/googlemrf.GoogleMRF/ValidateLOB',
          request_serializer=Query.SerializeToString,
          response_deserializer=Single.FromString,
          )
      self.ValidateMedia = channel.unary_unary(
          '/googlemrf.GoogleMRF/ValidateMedia',
          request_serializer=Query.SerializeToString,
          response_deserializer=Single.FromString,
          )
      self.ValidateProduct = channel.unary_unary(
          '/googlemrf.GoogleMRF/ValidateProduct',
          request_serializer=Query.SerializeToString,
          response_deserializer=Single.FromString,
          )
      self.ValidateSubMedia = channel.unary_unary(
          '/googlemrf.GoogleMRF/ValidateSubMedia',
          request_serializer=Query.SerializeToString,
          response_deserializer=Single.FromString,
          )
      self.ValidateSubProduct = channel.unary_unary(
          '/googlemrf.GoogleMRF/ValidateSubProduct',
          request_serializer=Query.SerializeToString,
          response_deserializer=Single.FromString,
          )


  class GoogleMRFServicer(object):
    """Interface exported by the server.
    """

    def MRFInfo(self, request, context):
      """Obtains a single MRF code and its related info.

      If it is not found an empty MRF is returned.
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def MRFs(self, request, context):
      """Obtains all MRF codes known.
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def Agencies(self, request, context):
      """Obtains all Agencies.
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def LOBs(self, request, context):
      """Obtains all LOBs known.
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def Products(self, request, context):
      """Obtains all Products known.
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def SubProducts(self, request, context):
      """Obtains all SubProducts known.
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def Channels(self, request, context):
      """Obtains all Channels known.
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def Medias(self, request, context):
      """Obtains all Medias known.
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def SubMedias(self, request, context):
      """Obtains all SubMedias known.
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def ValidateMRF(self, request, context):
      """The following methods check to see if the
      described object exists.

      In the case of something like a product, the query
      can be provided with both the name of the product
      and the parent line of business. If the combination
      of product and LOB does not exist empty will be
      returned.

      Returns an empty instance if it is not found.
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def ValidateAgency(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def ValidateChannel(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def ValidateLOB(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def ValidateMedia(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def ValidateProduct(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def ValidateSubMedia(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def ValidateSubProduct(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')


  def add_GoogleMRFServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'MRFInfo': grpc.unary_unary_rpc_method_handler(
            servicer.MRFInfo,
            request_deserializer=Query.FromString,
            response_serializer=MRF.SerializeToString,
        ),
        'MRFs': grpc.unary_stream_rpc_method_handler(
            servicer.MRFs,
            request_deserializer=Empty.FromString,
            response_serializer=MRF.SerializeToString,
        ),
        'Agencies': grpc.unary_stream_rpc_method_handler(
            servicer.Agencies,
            request_deserializer=Empty.FromString,
            response_serializer=Single.SerializeToString,
        ),
        'LOBs': grpc.unary_stream_rpc_method_handler(
            servicer.LOBs,
            request_deserializer=Empty.FromString,
            response_serializer=Single.SerializeToString,
        ),
        'Products': grpc.unary_stream_rpc_method_handler(
            servicer.Products,
            request_deserializer=Empty.FromString,
            response_serializer=Double.SerializeToString,
        ),
        'SubProducts': grpc.unary_stream_rpc_method_handler(
            servicer.SubProducts,
            request_deserializer=Empty.FromString,
            response_serializer=Double.SerializeToString,
        ),
        'Channels': grpc.unary_stream_rpc_method_handler(
            servicer.Channels,
            request_deserializer=Empty.FromString,
            response_serializer=Single.SerializeToString,
        ),
        'Medias': grpc.unary_stream_rpc_method_handler(
            servicer.Medias,
            request_deserializer=Empty.FromString,
            response_serializer=Double.SerializeToString,
        ),
        'SubMedias': grpc.unary_stream_rpc_method_handler(
            servicer.SubMedias,
            request_deserializer=Empty.FromString,
            response_serializer=Double.SerializeToString,
        ),
        'ValidateMRF': grpc.unary_unary_rpc_method_handler(
            servicer.ValidateMRF,
            request_deserializer=Query.FromString,
            response_serializer=MRF.SerializeToString,
        ),
        'ValidateAgency': grpc.unary_unary_rpc_method_handler(
            servicer.ValidateAgency,
            request_deserializer=Query.FromString,
            response_serializer=Single.SerializeToString,
        ),
        'ValidateChannel': grpc.unary_unary_rpc_method_handler(
            servicer.ValidateChannel,
            request_deserializer=Query.FromString,
            response_serializer=Single.SerializeToString,
        ),
        'ValidateLOB': grpc.unary_unary_rpc_method_handler(
            servicer.ValidateLOB,
            request_deserializer=Query.FromString,
            response_serializer=Single.SerializeToString,
        ),
        'ValidateMedia': grpc.unary_unary_rpc_method_handler(
            servicer.ValidateMedia,
            request_deserializer=Query.FromString,
            response_serializer=Single.SerializeToString,
        ),
        'ValidateProduct': grpc.unary_unary_rpc_method_handler(
            servicer.ValidateProduct,
            request_deserializer=Query.FromString,
            response_serializer=Single.SerializeToString,
        ),
        'ValidateSubMedia': grpc.unary_unary_rpc_method_handler(
            servicer.ValidateSubMedia,
            request_deserializer=Query.FromString,
            response_serializer=Single.SerializeToString,
        ),
        'ValidateSubProduct': grpc.unary_unary_rpc_method_handler(
            servicer.ValidateSubProduct,
            request_deserializer=Query.FromString,
            response_serializer=Single.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'googlemrf.GoogleMRF', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


  class BetaGoogleMRFServicer(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    """Interface exported by the server.
    """
    def MRFInfo(self, request, context):
      """Obtains a single MRF code and its related info.

      If it is not found an empty MRF is returned.
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def MRFs(self, request, context):
      """Obtains all MRF codes known.
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def Agencies(self, request, context):
      """Obtains all Agencies.
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def LOBs(self, request, context):
      """Obtains all LOBs known.
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def Products(self, request, context):
      """Obtains all Products known.
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def SubProducts(self, request, context):
      """Obtains all SubProducts known.
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def Channels(self, request, context):
      """Obtains all Channels known.
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def Medias(self, request, context):
      """Obtains all Medias known.
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def SubMedias(self, request, context):
      """Obtains all SubMedias known.
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def ValidateMRF(self, request, context):
      """The following methods check to see if the
      described object exists.

      In the case of something like a product, the query
      can be provided with both the name of the product
      and the parent line of business. If the combination
      of product and LOB does not exist empty will be
      returned.

      Returns an empty instance if it is not found.
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def ValidateAgency(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def ValidateChannel(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def ValidateLOB(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def ValidateMedia(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def ValidateProduct(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def ValidateSubMedia(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def ValidateSubProduct(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)


  class BetaGoogleMRFStub(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    """Interface exported by the server.
    """
    def MRFInfo(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """Obtains a single MRF code and its related info.

      If it is not found an empty MRF is returned.
      """
      raise NotImplementedError()
    MRFInfo.future = None
    def MRFs(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """Obtains all MRF codes known.
      """
      raise NotImplementedError()
    def Agencies(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """Obtains all Agencies.
      """
      raise NotImplementedError()
    def LOBs(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """Obtains all LOBs known.
      """
      raise NotImplementedError()
    def Products(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """Obtains all Products known.
      """
      raise NotImplementedError()
    def SubProducts(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """Obtains all SubProducts known.
      """
      raise NotImplementedError()
    def Channels(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """Obtains all Channels known.
      """
      raise NotImplementedError()
    def Medias(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """Obtains all Medias known.
      """
      raise NotImplementedError()
    def SubMedias(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """Obtains all SubMedias known.
      """
      raise NotImplementedError()
    def ValidateMRF(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """The following methods check to see if the
      described object exists.

      In the case of something like a product, the query
      can be provided with both the name of the product
      and the parent line of business. If the combination
      of product and LOB does not exist empty will be
      returned.

      Returns an empty instance if it is not found.
      """
      raise NotImplementedError()
    ValidateMRF.future = None
    def ValidateAgency(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    ValidateAgency.future = None
    def ValidateChannel(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    ValidateChannel.future = None
    def ValidateLOB(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    ValidateLOB.future = None
    def ValidateMedia(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    ValidateMedia.future = None
    def ValidateProduct(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    ValidateProduct.future = None
    def ValidateSubMedia(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    ValidateSubMedia.future = None
    def ValidateSubProduct(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    ValidateSubProduct.future = None


  def beta_create_GoogleMRF_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_deserializers = {
      ('googlemrf.GoogleMRF', 'Agencies'): Empty.FromString,
      ('googlemrf.GoogleMRF', 'Channels'): Empty.FromString,
      ('googlemrf.GoogleMRF', 'LOBs'): Empty.FromString,
      ('googlemrf.GoogleMRF', 'MRFInfo'): Query.FromString,
      ('googlemrf.GoogleMRF', 'MRFs'): Empty.FromString,
      ('googlemrf.GoogleMRF', 'Medias'): Empty.FromString,
      ('googlemrf.GoogleMRF', 'Products'): Empty.FromString,
      ('googlemrf.GoogleMRF', 'SubMedias'): Empty.FromString,
      ('googlemrf.GoogleMRF', 'SubProducts'): Empty.FromString,
      ('googlemrf.GoogleMRF', 'ValidateAgency'): Query.FromString,
      ('googlemrf.GoogleMRF', 'ValidateChannel'): Query.FromString,
      ('googlemrf.GoogleMRF', 'ValidateLOB'): Query.FromString,
      ('googlemrf.GoogleMRF', 'ValidateMRF'): Query.FromString,
      ('googlemrf.GoogleMRF', 'ValidateMedia'): Query.FromString,
      ('googlemrf.GoogleMRF', 'ValidateProduct'): Query.FromString,
      ('googlemrf.GoogleMRF', 'ValidateSubMedia'): Query.FromString,
      ('googlemrf.GoogleMRF', 'ValidateSubProduct'): Query.FromString,
    }
    response_serializers = {
      ('googlemrf.GoogleMRF', 'Agencies'): Single.SerializeToString,
      ('googlemrf.GoogleMRF', 'Channels'): Single.SerializeToString,
      ('googlemrf.GoogleMRF', 'LOBs'): Single.SerializeToString,
      ('googlemrf.GoogleMRF', 'MRFInfo'): MRF.SerializeToString,
      ('googlemrf.GoogleMRF', 'MRFs'): MRF.SerializeToString,
      ('googlemrf.GoogleMRF', 'Medias'): Double.SerializeToString,
      ('googlemrf.GoogleMRF', 'Products'): Double.SerializeToString,
      ('googlemrf.GoogleMRF', 'SubMedias'): Double.SerializeToString,
      ('googlemrf.GoogleMRF', 'SubProducts'): Double.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateAgency'): Single.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateChannel'): Single.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateLOB'): Single.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateMRF'): MRF.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateMedia'): Single.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateProduct'): Single.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateSubMedia'): Single.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateSubProduct'): Single.SerializeToString,
    }
    method_implementations = {
      ('googlemrf.GoogleMRF', 'Agencies'): face_utilities.unary_stream_inline(servicer.Agencies),
      ('googlemrf.GoogleMRF', 'Channels'): face_utilities.unary_stream_inline(servicer.Channels),
      ('googlemrf.GoogleMRF', 'LOBs'): face_utilities.unary_stream_inline(servicer.LOBs),
      ('googlemrf.GoogleMRF', 'MRFInfo'): face_utilities.unary_unary_inline(servicer.MRFInfo),
      ('googlemrf.GoogleMRF', 'MRFs'): face_utilities.unary_stream_inline(servicer.MRFs),
      ('googlemrf.GoogleMRF', 'Medias'): face_utilities.unary_stream_inline(servicer.Medias),
      ('googlemrf.GoogleMRF', 'Products'): face_utilities.unary_stream_inline(servicer.Products),
      ('googlemrf.GoogleMRF', 'SubMedias'): face_utilities.unary_stream_inline(servicer.SubMedias),
      ('googlemrf.GoogleMRF', 'SubProducts'): face_utilities.unary_stream_inline(servicer.SubProducts),
      ('googlemrf.GoogleMRF', 'ValidateAgency'): face_utilities.unary_unary_inline(servicer.ValidateAgency),
      ('googlemrf.GoogleMRF', 'ValidateChannel'): face_utilities.unary_unary_inline(servicer.ValidateChannel),
      ('googlemrf.GoogleMRF', 'ValidateLOB'): face_utilities.unary_unary_inline(servicer.ValidateLOB),
      ('googlemrf.GoogleMRF', 'ValidateMRF'): face_utilities.unary_unary_inline(servicer.ValidateMRF),
      ('googlemrf.GoogleMRF', 'ValidateMedia'): face_utilities.unary_unary_inline(servicer.ValidateMedia),
      ('googlemrf.GoogleMRF', 'ValidateProduct'): face_utilities.unary_unary_inline(servicer.ValidateProduct),
      ('googlemrf.GoogleMRF', 'ValidateSubMedia'): face_utilities.unary_unary_inline(servicer.ValidateSubMedia),
      ('googlemrf.GoogleMRF', 'ValidateSubProduct'): face_utilities.unary_unary_inline(servicer.ValidateSubProduct),
    }
    server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
    return beta_implementations.server(method_implementations, options=server_options)


  def beta_create_GoogleMRF_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_serializers = {
      ('googlemrf.GoogleMRF', 'Agencies'): Empty.SerializeToString,
      ('googlemrf.GoogleMRF', 'Channels'): Empty.SerializeToString,
      ('googlemrf.GoogleMRF', 'LOBs'): Empty.SerializeToString,
      ('googlemrf.GoogleMRF', 'MRFInfo'): Query.SerializeToString,
      ('googlemrf.GoogleMRF', 'MRFs'): Empty.SerializeToString,
      ('googlemrf.GoogleMRF', 'Medias'): Empty.SerializeToString,
      ('googlemrf.GoogleMRF', 'Products'): Empty.SerializeToString,
      ('googlemrf.GoogleMRF', 'SubMedias'): Empty.SerializeToString,
      ('googlemrf.GoogleMRF', 'SubProducts'): Empty.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateAgency'): Query.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateChannel'): Query.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateLOB'): Query.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateMRF'): Query.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateMedia'): Query.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateProduct'): Query.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateSubMedia'): Query.SerializeToString,
      ('googlemrf.GoogleMRF', 'ValidateSubProduct'): Query.SerializeToString,
    }
    response_deserializers = {
      ('googlemrf.GoogleMRF', 'Agencies'): Single.FromString,
      ('googlemrf.GoogleMRF', 'Channels'): Single.FromString,
      ('googlemrf.GoogleMRF', 'LOBs'): Single.FromString,
      ('googlemrf.GoogleMRF', 'MRFInfo'): MRF.FromString,
      ('googlemrf.GoogleMRF', 'MRFs'): MRF.FromString,
      ('googlemrf.GoogleMRF', 'Medias'): Double.FromString,
      ('googlemrf.GoogleMRF', 'Products'): Double.FromString,
      ('googlemrf.GoogleMRF', 'SubMedias'): Double.FromString,
      ('googlemrf.GoogleMRF', 'SubProducts'): Double.FromString,
      ('googlemrf.GoogleMRF', 'ValidateAgency'): Single.FromString,
      ('googlemrf.GoogleMRF', 'ValidateChannel'): Single.FromString,
      ('googlemrf.GoogleMRF', 'ValidateLOB'): Single.FromString,
      ('googlemrf.GoogleMRF', 'ValidateMRF'): MRF.FromString,
      ('googlemrf.GoogleMRF', 'ValidateMedia'): Single.FromString,
      ('googlemrf.GoogleMRF', 'ValidateProduct'): Single.FromString,
      ('googlemrf.GoogleMRF', 'ValidateSubMedia'): Single.FromString,
      ('googlemrf.GoogleMRF', 'ValidateSubProduct'): Single.FromString,
    }
    cardinalities = {
      'Agencies': cardinality.Cardinality.UNARY_STREAM,
      'Channels': cardinality.Cardinality.UNARY_STREAM,
      'LOBs': cardinality.Cardinality.UNARY_STREAM,
      'MRFInfo': cardinality.Cardinality.UNARY_UNARY,
      'MRFs': cardinality.Cardinality.UNARY_STREAM,
      'Medias': cardinality.Cardinality.UNARY_STREAM,
      'Products': cardinality.Cardinality.UNARY_STREAM,
      'SubMedias': cardinality.Cardinality.UNARY_STREAM,
      'SubProducts': cardinality.Cardinality.UNARY_STREAM,
      'ValidateAgency': cardinality.Cardinality.UNARY_UNARY,
      'ValidateChannel': cardinality.Cardinality.UNARY_UNARY,
      'ValidateLOB': cardinality.Cardinality.UNARY_UNARY,
      'ValidateMRF': cardinality.Cardinality.UNARY_UNARY,
      'ValidateMedia': cardinality.Cardinality.UNARY_UNARY,
      'ValidateProduct': cardinality.Cardinality.UNARY_UNARY,
      'ValidateSubMedia': cardinality.Cardinality.UNARY_UNARY,
      'ValidateSubProduct': cardinality.Cardinality.UNARY_UNARY,
    }
    stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
    return beta_implementations.dynamic_stub(channel, 'googlemrf.GoogleMRF', cardinalities, options=stub_options)
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)
