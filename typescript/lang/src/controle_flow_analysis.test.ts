/**
 * https://www.typescriptlang.org/static/TypeScript%20Control%20Flow%20Analysis-8a549253ad8470850b77c4c5c351d457.png
 */

describe('if文', () => {
  it('if-else', () => {

  });
  it('typeof', () => {
    const input: string | number = 'who am i';
    if (typeof input === 'string') {
      expect(typeof input === 'string').toBeTruthy();
      expect(typeof input === 'number').toBeFalsy();
    }
  });
  it('instanceof', () => {
    const input: number | number[] = [1, 2, 3];
    if (input instanceof Array) {
      expect(input[0]).toBe(1); // 配列としてアクセス可能
    }
  });
  it('"property" in object', () => {
    interface Foo {
      foo: string;
    };
    const input: Foo = {
      foo: 'foobar',
    };
    if ('foo' in input) {
      expect(input.foo).toBe('foobar'); // fooプロパティにアクセス可能
    }
  });
});
