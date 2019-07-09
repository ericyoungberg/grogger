# Grogger
### Pass the `console.log` around

For logging across browsers in pub-sub fashion, using nothing more than a `<script>` tag.


## Installation
#### Get the files
```
git clone https://github.com/ericyoungberg/grogger
cd grogger
```

#### With Go (linux/amd64)
```
make
sudo make install
```


## Usage

```
grogger [-p portNumber]
```

#### Browser

Now you can add a `<script>` tag for the host where the Grogger instance is located with an associated session endpoint

```html
<head>
    ...
    <script src="https://logs.ericyoungberg.com/some-session"></script>
    ...
</head>
```
