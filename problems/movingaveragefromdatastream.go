package main

type MovingAverage struct {
	size  int
	count int
	deque []float64
	sum   float64
}

func Constructor13(size int) MovingAverage {
	return MovingAverage{
		size:  size,
		deque: make([]float64, 0),
		sum:   0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.sum += float64(val)
	this.deque = append(this.deque, float64(val))

	if len(this.deque) > this.size {
		this.sum -= this.deque[0]
		this.deque = this.deque[1:]
	}

	return this.sum / float64(len(this.deque))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
