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

        console[method] = function(...args) {
            if (typeof args[args.length - 1] === 'object' && args[args.length - 1].__preventGrog) {
                args.pop();
            }
            else {
                socket.send(JSON.stringify({
                    message: args,
                    method,
                }));
            }

            _method(...args);
        };
    });

    socket.addEventListener('open', e => {
        console.info('Joined the session');
    });

    socket.addEventListener('message', ({ data }) => {
        data = JSON.parse(data);
        console[data.method](data.sender, ': ', ...data.message, { __preventGrog: true });
    });

})();
