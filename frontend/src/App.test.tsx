import { describe, it, expect } from 'vitest';
import {sayHello} from "./App";

describe('something truthy and falsy', () => {
    it('to say hello', () => {
        expect(sayHello()).toBe('Hello, World!');
    });
});
