class ByteUtils {
    static serialize(obj) {
        var object = {dd:"ddd", sub:{xx:"dd"}, num:666};
        var string = JSON.stringify(object);
        var uint8_array = new TextEncoder(document.characterSet.toLowerCase()).encode(string);
        var array_buffer = uint8_array.buffer;

        return array_buffer
    }

    static deserialize(array_buffer) {
        // Now to the decoding
        var decoder = new TextDecoder("utf-8");
        var view = new DataView(array_buffer, 0, array_buffer.byteLength);
        var string = decoder.decode(view);
        var object = JSON.parse(string);

        return object
    }
}
