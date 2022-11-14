package software.amazon.dafny.conversion;

import java.math.BigInteger;
import java.nio.ByteBuffer;
import java.time.Instant;
import java.util.Date;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.Function;

import dafny.Array;
import dafny.DafnyMap;
import dafny.DafnySequence;
import dafny.DafnySet;
import dafny.Tuple2;
import dafny.TypeDescriptor;


/**
 * Methods that convert from a Native Java type to a Dafny Java type.
 */
public class ToDafny {

    /**
     * Methods that convert from a Native Java type
     * to a Dafny Java type,
     * for Smithy's definition of Simple shapes.
     */
    public static class Simple {

        // BLOB("blob", BlobShape.class, Category.SIMPLE),
        public static DafnySequence<Byte> ByteSequence(ByteBuffer byteBuffer) {
            return DafnySequence.fromArray(TypeDescriptor.BYTE, Array.wrap(byteBuffer.array()));
        }

        // STRING("string", StringShape.class, Category.SIMPLE),
        public static DafnySequence<Character> CharacterSequence(String s) {
            return DafnySequence.asString(s);
        }

        // TIMESTAMP("timestamp", TimestampShape.class, Category.SIMPLE),
        public static DafnySequence<Character> CharacterSequence(Date input) {
            // KMS uses unix timestamp, or seconds from epoch, as its serialized timestamp
            // Other services may use other formats
            return CharacterSequence(String.format("%d", (input.getTime() / 1000L)));
        }

        // TIMESTAMP("timestamp", TimestampShape.class, Category.SIMPLE),
        public static DafnySequence<Character> CharacterSequence(Instant input) {
            // KMS uses unix timestamp, or seconds from epoch, as its serialized timestamp
            // Other services may use other formats
            return CharacterSequence(String.format("%d", input.getEpochSecond()));
        }

    }

    public static class Aggregate {

        // LIST("list", ListShape.class, Category.AGGREGATE),
        public static <INPUT, OUTPUT> DafnySequence<? extends OUTPUT> GenericToSequence(
                List<INPUT> nativeValues,
                Function<INPUT, OUTPUT> converter,
                TypeDescriptor<OUTPUT> typeDescriptor
        ) {
            class Local {
                public OUTPUT IndexConverter(BigInteger index) {
                    return converter.apply(nativeValues.get(index.intValue()));
                }
            }
            Local local = new Local();
            return DafnySequence.Create(
                    typeDescriptor,
                    BigInteger.valueOf(nativeValues.size()),
                    local::IndexConverter);
        }

        // SET("set", SetShape.class, Category.AGGREGATE),
        // TODO: Frankly, we should avoid Dafny Sets since they do not preserve order
        // But, we would need to implement our own Dafny Ordered Set...
        public static <INPUT, OUTPUT> DafnySet<OUTPUT> GenericToSet(
                Set<INPUT> nativeValues,
                Function<INPUT, OUTPUT> converter
        ) {
            HashSet<OUTPUT> hashSet = new HashSet<>(nativeValues.size(), 1);
            nativeValues.forEach(v -> hashSet.add(converter.apply(v)));
            return new DafnySet<>(hashSet);
        }

        // MAP("map", MapShape.class, Category.AGGREGATE),
        // Technically, a smithy Map's Key value will always be a String
        public static <IN_KEY, IN_VALUE, OUT_KEY, OUT_VALUE> DafnyMap<OUT_KEY, OUT_VALUE> GenericToMap(
                Map<IN_KEY, IN_VALUE> nativeValue,
                Function<IN_KEY, OUT_KEY> keyConverter,
                Function<IN_VALUE, OUT_VALUE> valueConverter
        ) {
            @SuppressWarnings("unchecked")
            Tuple2<OUT_KEY, OUT_VALUE>[] tuples = new Tuple2[nativeValue.size()];
            AtomicInteger counter = new AtomicInteger(0);
            nativeValue.forEach((k, v) -> tuples[counter.getAndIncrement()] = new Tuple2<>(
                    keyConverter.apply(k), valueConverter.apply(v))
            );
            return DafnyMap.fromElements(tuples);
        }
    }
}