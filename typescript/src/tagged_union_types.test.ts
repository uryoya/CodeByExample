describe('タグ付きユニオン型 (Tagged Union Type / Discriminated Union Type)', () => {
  type Result<T> = Ok<T> | Failed;
  type Ok<T> = {
    type: 'ok';
    value: T;
  }
  type Failed = {
    type: 'failed';
    error: Error;
  }
  const ok: Result<string> = {
    type: 'ok',
    value: 'ok value',
  };
  const error: Result<string> = {
    type: 'failed',
    error: new Error(),
  };
  it('共通して持つtypeプロパティを使って型の絞り込みができる (if version)', () => {
    function unwrapResult<T>(result: Result<T>): T {
      if (result.type === 'ok') {
        return result.value;
      }
      // 型の絞り込みができているのでエラーにならない
      throw result.error;
    }
    expect(unwrapResult(ok)).toBe('ok value');
    expect(() => unwrapResult(error)).toThrow(Error);
  });
  it('共通して持つtypeプロパティを使って型の絞り込みができる (switch version)', () => {
    function unwrapResult<T>(result: Result<T>): T {
      switch (result.type) {
        case 'ok':
          return result.value;
        case 'failed':
          // 型の絞り込みができているのでエラーにならない
          throw result.error;
      }
    }
    expect(unwrapResult(ok)).toBe('ok value');
    expect(() => unwrapResult(error)).toThrow(Error);
  });
  it('タグの値を事前に変数に入れることができる', () => {
    function unwrapResult<T>(result: Result<T>): T {
      const { type } = result; // 事前に変数に入れる
      switch (type) {
        // ちゃんと絞り込みができている
        case 'ok':
          return result.value;
        case 'failed':
          throw result.error;
      }
    }
    expect(unwrapResult(ok)).toBe('ok value');
    expect(() => unwrapResult(error)).toThrow(Error);
  });
  it('タグに対する判別式を変数に入れることができる', () => {
    function unwrapResult<T>(result: Result<T>): T {
      const isOk = result.type === 'ok';
      if (isOk) {
        // ちゃんと絞り込みができている
        return result.value;
      }
      throw result.error;
    }
    expect(unwrapResult(ok)).toBe('ok value');
    expect(() => unwrapResult(error)).toThrow(Error);
  })
});
