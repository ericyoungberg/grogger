(() => {

    const socket = new WebSocket('ws://{{SESSION}}');
    const _console = console;

    const methods = [
        'error',
        'info',
        'log',
        'warn',
    ];

    methods.forEach(method => {
        const _method = _console[method];

        console[method] = (message, send = true) => {
            _method(message);

            if (send) {
                socket.send(JSON.stringify({
                    message,
                    method,
                }));
            }
        };
    });

    socket.addEventListener('open', e => {
        console.info('Joined the session')
    });

    socket.addEventListener('message', ({ data }) => {
        console.log(data, false);
    });

})();
