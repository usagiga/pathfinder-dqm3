# Monsters

Generating from https://game8.jp/dqm3/552459

To generate monster list,
paste and execute these code on browser console.

```js
const parseFamily = (name) => {
    if (name === 'スライム系') {
        return 'Slime';
    }
    if (name === 'ドラゴン系') {
        return 'Dragon';
    }
    if (name === '自然系') {
        return 'Nature';
    }
    if (name === '魔獣系') {
        return 'Beast';
    }
    if (name === '物質系') {
        return 'Material';
    }
    if (name === '悪魔系') {
        return 'Demon';
    }
    if (name === 'ゾンビ系') {
        return 'Undead';
    }
    return 'Unknown';
};
```

```js
let monsters = [...document.querySelectorAll('._1My64rRAozhRV0eyxESiJN')];
```

```js
let monstersObj = monsters.map(m => { return {
    ID: m.querySelector('.EhZqLMa9S8wHQWJSecnd9').textContent.replace('No.', ''),
    Name: m.querySelector('.O4LptJfHPAXWXEDKLmV7I').textContent,
    Family: 'common.Family_' + parseFamily(m.querySelector('._2M4m7hs_RdJb674LYgZX6f').textContent),
    Rank: 'common.Rank_' + m.querySelector('._3S10RRo8USXJVRhobC6pcN').textContent,
}})
```

```js
monsterObjs.map(m => '{ID: ' + m.ID + ', Name: "' + m.Name + '", Family: ' + m.Family + ', Rank: ' + m.Rank + '}')
```

then copy its result!
