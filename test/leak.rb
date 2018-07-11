require 'objspace'

ObjectSpace.trace_object_allocations_start

class Leaky
  def self.leak
    @leak ||= []
  end

  def doit
    noleak = []

    100.times do
      self.class.leak << 'leaked memory'
      noleak << 'not leaked'
    end
  end
end

def dump_heap(filename)
  GC.start
  File.open(filename, 'w') { |file| ObjectSpace.dump_all(output: file) }
end

Leaky.new.doit
dump_heap(File.expand_path('./heap1.jsonl', __dir__))
Leaky.new.doit
dump_heap(File.expand_path('./heap2.jsonl', __dir__))
Leaky.new.doit
dump_heap(File.expand_path('./heap3.jsonl', __dir__))
