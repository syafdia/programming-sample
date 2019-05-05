
class Queue():

    def __init__(self):
        self._xs = []

    def __str__(self):
        return ','.join(map(str, self._xs))

    def size(self):
        return len(self._xs)

    def enqueue(self, x):
        self._xs.append(x)

    def dequeue(self):
        if self.size() == 0:
            return None
        
        return self._xs.pop(0)


class Stack:
    
    def __init__(self):
        self._q = Queue()

    def __str__(self):
        return '{}'.format(self._q)

    def push(self, v):
        self._q.enqueue(v)

    def pop(self):
        if self._q.size() == 0:
            return None

        tmp_q = Queue()

        while self._q.size() > 1:
            tmp_q.enqueue(self._q.dequeue())

        top = self._q.dequeue()
        self._q = tmp_q

        return top



def run_queue():
    queue = Queue()
    queue.enqueue(1)
    queue.enqueue(2)
    print(queue)
    print(queue.dequeue())
    print(queue.dequeue())
    print(queue)

def run_stack():
    stack = Stack()
    stack.push(1)
    stack.push(2)
    stack.push(3)
    stack.push(4)
    print(stack)
    print(stack.pop())
    print(stack.pop())
    print(stack)

if __name__ == "__main__":
    run_queue()
    run_stack()
