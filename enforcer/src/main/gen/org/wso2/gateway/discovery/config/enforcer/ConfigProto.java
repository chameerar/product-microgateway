// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/config/enforcer/config.proto

package org.wso2.gateway.discovery.config.enforcer;

public final class ConfigProto {
  private ConfigProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_wso2_discovery_config_enforcer_Config_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_wso2_discovery_config_enforcer_Config_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n+wso2/discovery/config/enforcer/config." +
      "proto\022\036wso2.discovery.config.enforcer\032)w" +
      "so2/discovery/config/enforcer/cert.proto" +
      "\032+wso2/discovery/config/enforcer/issuer." +
      "proto\032.wso2/discovery/config/enforcer/ev" +
      "ent_hub.proto\0323wso2/discovery/config/enf" +
      "orcer/am_credentials.proto\0321wso2/discove" +
      "ry/config/enforcer/auth_service.proto\0322w" +
      "so2/discovery/config/enforcer/jwt_genera" +
      "tor.proto\032/wso2/discovery/config/enforce" +
      "r/throttling.proto\"\224\004\n\006Config\022>\n\016jwtToke" +
      "nConfig\030\001 \003(\0132&.wso2.discovery.config.en" +
      "forcer.Issuer\022;\n\010keystore\030\002 \001(\0132).wso2.d" +
      "iscovery.config.enforcer.CertStore\022=\n\ntr" +
      "uststore\030\003 \001(\0132).wso2.discovery.config.e" +
      "nforcer.CertStore\022:\n\010eventhub\030\004 \001(\0132(.ws" +
      "o2.discovery.config.enforcer.EventHub\022@\n" +
      "\013authService\030\005 \001(\0132+.wso2.discovery.conf" +
      "ig.enforcer.AuthService\022F\n\017apimCredentia" +
      "ls\030\006 \001(\0132-.wso2.discovery.config.enforce" +
      "r.AmCredentials\022B\n\014jwtGenerator\030\007 \001(\0132,." +
      "wso2.discovery.config.enforcer.JWTGenera" +
      "tor\022D\n\020throttlingConfig\030\010 \001(\0132*.wso2.dis" +
      "covery.config.enforcer.ThrottlingB\213\001\n*or" +
      "g.wso2.gateway.discovery.config.enforcer" +
      "B\013ConfigProtoP\001ZNgithub.com/envoyproxy/g" +
      "o-control-plane/wso2/discovery/config/en" +
      "forcer;enforcerb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.wso2.gateway.discovery.config.enforcer.CertStoreProto.getDescriptor(),
          org.wso2.gateway.discovery.config.enforcer.IssuerProto.getDescriptor(),
          org.wso2.gateway.discovery.config.enforcer.EventHubProto.getDescriptor(),
          org.wso2.gateway.discovery.config.enforcer.AmCredentialsProto.getDescriptor(),
          org.wso2.gateway.discovery.config.enforcer.AuthServiceProto.getDescriptor(),
          org.wso2.gateway.discovery.config.enforcer.JWTGeneratorProto.getDescriptor(),
          org.wso2.gateway.discovery.config.enforcer.ThrottlingProto.getDescriptor(),
        });
    internal_static_wso2_discovery_config_enforcer_Config_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_wso2_discovery_config_enforcer_Config_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_wso2_discovery_config_enforcer_Config_descriptor,
        new java.lang.String[] { "JwtTokenConfig", "Keystore", "Truststore", "Eventhub", "AuthService", "ApimCredentials", "JwtGenerator", "ThrottlingConfig", });
    org.wso2.gateway.discovery.config.enforcer.CertStoreProto.getDescriptor();
    org.wso2.gateway.discovery.config.enforcer.IssuerProto.getDescriptor();
    org.wso2.gateway.discovery.config.enforcer.EventHubProto.getDescriptor();
    org.wso2.gateway.discovery.config.enforcer.AmCredentialsProto.getDescriptor();
    org.wso2.gateway.discovery.config.enforcer.AuthServiceProto.getDescriptor();
    org.wso2.gateway.discovery.config.enforcer.JWTGeneratorProto.getDescriptor();
    org.wso2.gateway.discovery.config.enforcer.ThrottlingProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
