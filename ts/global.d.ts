declare type BinaryNode<T> = {
	left: BinaryNode<T> | null;
	right: BinaryNode<T> | null;
	value: T;
}

