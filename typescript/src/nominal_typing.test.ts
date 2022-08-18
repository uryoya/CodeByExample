// TypeScriptは構造的部分型(Structural Subtyping)を採用しているため、型のシグネチャが同じであれば同じ型とみなされる。
// 柔軟性がある一方、構造が同じでも名前が違う型を区別(=これを公称型/Nominal Typeという)したい時に不便。そこでTypeScript
// でも擬似的に公称型を使うためのテクニックが以下のページで紹介されている。
// https://speakerdeck.com/okunokentaro/harajukuts2
// https://typescript-jp.gitbook.io/deep-dive/main-1/nominaltyping
// 1つめのページで紹介されている内容を試す。

// 公称型の基本形を定義
// { __brand: U } というオブジェクトとの交差型にすることでリテラルのまま区別できるようにする（原文ママ）
type Nominal<T, U extends string> = T & { __brand: U };

describe('Nominalを使ってIDを区別する', () => {
  it('type aliasだと型の区別ができない', () => {
    type BookId = string;
    type UserId = string;
    function fetchBook(id: BookId) { console.log(id) }

    const id = 'cSh8sJrD3AOVBaZrpiFT' as UserId;
    fetchBook(id); // トランスパイルエラーにならない
  });
  it('型の区別ができている', () => {
    type BookId = Nominal<string, 'BookId'>;
    type UserId = Nominal<string, 'UserId'>;
    function fetchBook(id: BookId) { console.log(id); }
    function fetchUser(id: UserId) { console.log(id) }

    const userId = 'cSh8sJrD3AOVBaZrpiFT' as UserId;
    const bookId = 'j300jd9f2ojjjdjwfwoD' as BookId;
    // fetchBook(userId); // トランスパイルエラーになる
    fetchUser(userId);
    fetchBook(bookId);
  });
  it('asを使わずに値を作る', () => {
    type NonEmptyString = Nominal<string, 'NonEmptyString'>;
    // NonEmptyStringかどうか検証する関数
    function isNonEmptyString(v: unknown): v is NonEmptyString {
      return typeof v === 'string' && v.length > 1;
    }
    function assertNonEmptyString(v: unknown, target = ''): asserts v is NonEmptyString {
      if (!isNonEmptyString(v)) {
        throw new Error(`${target} should be non empty string`.trim());
      }
    }
    // NonEmptyStringを使う関数
    function print(s: NonEmptyString) { console.log(s) }
    const str = 'foo';
    // print(str); // エラーになる
    assertNonEmptyString(str);
    print(str); // エラーにならない
    expect(isNonEmptyString('')).toBeFalsy();
  });
});
