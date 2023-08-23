// PLEASE IGNORE - WORK IN PROGRESS

package org.outline.sdk.demo.connectivity.capacitor

import com.getcapacitor.JSObject;
import com.getcapacitor.Plugin;
import com.getcapacitor.PluginCall;
import com.getcapacitor.PluginMethod;
import com.getcapacitor.annotation.CapacitorPlugin;
import com.myapp.sharedbackend.SharedBackend;
import kotlinx.serialization.*;
import kotlinx.serialization.json.JSON;

@Serializable
data class RawBackendCallInputMessage(
  val method: String,
  val input: String
)

@Serializable
data class RawBackendCallOutputMessage(
  val result: String,
  val errors: List<String>
)

@CapacitorPlugin(name = "MobileBackend")
class MobileBackendPlugin: Plugin() {
  @PluginMethod()
  fun Invoke(call: PluginCall) {
    val output = new JSObject();
    val outputMessage: RawBackendCallOutputMessage;

    try {
      val rawInputMessage = JSON.stringify(
        RawBackendCallInputMessage.serializer(),
          SharedBackend.sendRawCall(RawBackendCallInputMessage(
            call.getString("method")!!,
            call.getString("input")!!
          )
        )
      );

      outputMessage = JSON.parse(
        RawBackendCallOutputMessage.serializer(),
        rawInputMessage
      );
    } catch (error: Exception) {
      output.put("errors", listOf(error.message));

      return call.resolve(output);
    }

    output.put("result", outputMessage.result);
    output.put("errors", outputMessage.errors);

    return call.resolve(output);
  }
}