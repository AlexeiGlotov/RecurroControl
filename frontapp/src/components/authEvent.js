import EventEmitter from 'events';

class AuthEventEmitter extends EventEmitter {}
const authEvent = new AuthEventEmitter();

export default authEvent;