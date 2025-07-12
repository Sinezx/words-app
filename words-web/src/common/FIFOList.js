class FIFOListNode{
    constructor(data){
        this.data = data
        this.next = null
    }
}

export class FIFOList{
    constructor() {
        this.head = null
        this.tail = null
    }

    add(data){
        let node = new FIFOListNode(data)
        if(this.head == null){
            this.head = node
            this.tail = node
        }else{
            this.tail.next = node
            this.tail = node
        }
    }

    get(){
        if(this.head == null){
            return this.head
        }else{
            let node = this.head
            this.head = node.next
            if(node == this.tail){
                this.tail = null
            }
            return node.data
        }
    }

    isEmpty(){
        return this.head == null
    }
}