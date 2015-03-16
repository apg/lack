**lack**: A tool for querying logfmt style messages.

## Note:

This is super hacky right now, but "works." 

## Build

```
$ go generate ./... && go build ./... && go build
```

## Run

```
$ lack -format "<fmt string>" <query> < INPUT
```

## Format strings

### Positional (always ' ' delimited for now)

```
$ echo "foo bar" | ./lack -format '%1 %2 %0' 'foo'
bar foo foo bar
```

### Keys

```
$ echo "foo=bar" | ./lack -format '%{foo}' 'foo'
bar
```

### Mixed

```
$ echo "foo=bar baz" | ./lack -format '%{foo} %2' 'foo'
bar baz
```

### Inserted tabs (Note: inserted are ASCII 9, e.g. '\t', not spaces)

```
$ echo "foo=bar baz" | ./lack -format '%{foo}\t%2' 'foo'
bar     baz
```

## Current supported query types:

### Single word "grep"

```
$ echo "foo" | ./lack foo
foo
```

### Phrase "grep"

```
$ echo "foo bar" | ./lack '"foo bar"'
foo bar
```

### Regexp "grep"

```
$ echo "foo bar" | ./lack /fo+/
foo bar
```

### Field extraction

```
$ echo "foo=hello" | ./lack foo=hello
foo=hello
```

```
$ echo "foo=hello" | ./lack foo!=bar
foo=hello
```

```
$ echo "foo=6" | ./lack foo!=6
```

```
$ echo "foo=6" | ./lack 'foo>=6'
foo=6
```

```
$ echo "foo=6" | ./lack 'foo<6'
```

```
$ echo "foo=6" | ./lack 'foo<=6'
foo=6
```

```
$ echo "foo=bar" | ./lack 'foo=/bar/'
foo=bar
```

```
$ echo "foo=baz" | ./lack 'foo!=/bar/'
foo=baz
```

```
$ echo "foo=1" | ./lack 'foo!=0'
foo=1
```

### Conjunctions

#### And

```
$ echo "foo=6 bar=10" | ./lack 'foo>5 & bar<15'
foo=6 bar=10
```

#### Or

```
$ echo "foo=5 bar=15" | ./lack 'foo=5 | bar=20'
foo=5 bar=15
```
