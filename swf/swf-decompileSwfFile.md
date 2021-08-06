# Decompile a SWF file

A SWF file, while it can be played, can also be decompiled to view assets and scripts within it. One such tool to do this is to use [jpexs-decompiler](https://github.com/jindrapetrik/jpexs-decompiler/).

## Requirements

* Java 8+ installed

## Extract scripts

```
java -jar ffdec.jar -export script "/home/akp/swf/decompiled" myfile.swf
```

## Run interactive GUI

```
java -jar ffdec.jar myfile.swf
```
