const ffi = require('ffi-napi');

const libm = ffi.Library('libm', {
  ceil: ['double', ['double']]
});
console.log('ffi', libm.ceil(1.5)); // 2

const current = ffi.Library(null, {
  'atoi': ['int', ['string']]
});
console.log(current.atoi('123')); // 123
