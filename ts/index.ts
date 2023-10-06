//linked lists

type Node<T> = {
	value: T,
	next?: Node<T>,
}

export class Queue<T> {
	public length: number;
	private head?: Node<T> | undefined;
	private tail?: Node<T> | undefined;

	constructor() {
		this.head = this.tail = undefined;
		this.length = 0;
	}

	enqueue(item: T): void {
		if (!this.tail) {
			this.tail = this.head = {value: item} as Node<T>;
		}

		const node = {value: item} as Node<T>;
		this.tail.next = node;
		this.tail = node;
	}

	deque(): T | undefined {
		if (!this.head) {
			return undefined;
		}

		this.length--;
		const head = this.head;
		this.head = this.head.next;
		head.next = undefined;

		return head.value;
	}

	peek(): T | undefined {
		return this.head?.value;
	}
}

const q = new Queue();
q.enqueue(12);
q.enqueue(34);
q.enqueue(55);

