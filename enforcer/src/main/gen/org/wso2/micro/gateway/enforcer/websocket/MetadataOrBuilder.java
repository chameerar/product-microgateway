// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mgw_websocket_rls.proto

package org.wso2.micro.gateway.enforcer.websocket;

public interface MetadataOrBuilder extends
    // @@protoc_insertion_point(interface_extends:envoy.extensions.filters.http.mgw_websocket.v3.Metadata)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * Key is the reverse DNS filter name, e.g. com.acme.widget. The envoy.*
   * namespace is reserved for Envoy's built-in filters.
   * </pre>
   *
   * <code>map&lt;string, .google.protobuf.Struct&gt; filter_metadata = 1;</code>
   */
  int getFilterMetadataCount();
  /**
   * <pre>
   * Key is the reverse DNS filter name, e.g. com.acme.widget. The envoy.*
   * namespace is reserved for Envoy's built-in filters.
   * </pre>
   *
   * <code>map&lt;string, .google.protobuf.Struct&gt; filter_metadata = 1;</code>
   */
  boolean containsFilterMetadata(
      java.lang.String key);
  /**
   * Use {@link #getFilterMetadataMap()} instead.
   */
  @java.lang.Deprecated
  java.util.Map<java.lang.String, com.google.protobuf.Struct>
  getFilterMetadata();
  /**
   * <pre>
   * Key is the reverse DNS filter name, e.g. com.acme.widget. The envoy.*
   * namespace is reserved for Envoy's built-in filters.
   * </pre>
   *
   * <code>map&lt;string, .google.protobuf.Struct&gt; filter_metadata = 1;</code>
   */
  java.util.Map<java.lang.String, com.google.protobuf.Struct>
  getFilterMetadataMap();
  /**
   * <pre>
   * Key is the reverse DNS filter name, e.g. com.acme.widget. The envoy.*
   * namespace is reserved for Envoy's built-in filters.
   * </pre>
   *
   * <code>map&lt;string, .google.protobuf.Struct&gt; filter_metadata = 1;</code>
   */

  com.google.protobuf.Struct getFilterMetadataOrDefault(
      java.lang.String key,
      com.google.protobuf.Struct defaultValue);
  /**
   * <pre>
   * Key is the reverse DNS filter name, e.g. com.acme.widget. The envoy.*
   * namespace is reserved for Envoy's built-in filters.
   * </pre>
   *
   * <code>map&lt;string, .google.protobuf.Struct&gt; filter_metadata = 1;</code>
   */

  com.google.protobuf.Struct getFilterMetadataOrThrow(
      java.lang.String key);
}
