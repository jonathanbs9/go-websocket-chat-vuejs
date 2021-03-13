const socket = io();

new Vue({
    el: '#chat-app',
    created(){
        const  vm = this;
        socket.on("chat message",  (msg) => {
            vm.messages.push({
                text: msg,
                date: new Date().toLocaleDateString()
            })
        })
    },
    data: {
        message : '',
        messages : [{
            text: 'Hello this is a message!!!!!!!',
            date: new Date()
        }]
    },
    methods: {
        sendMessage(){
            socket.emit('chat message', this.message);
            this.message = '';
        }
    }
});
