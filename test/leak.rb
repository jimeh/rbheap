require 'objspace'

class Leaky
  def self.leak
    @leak ||= []
  end

  def doit
    noleak = []

    50.times do
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
dump_heap('heap1.jsonl')
Leaky.new.doit
dump_heap('heap2.jsonl')
Leaky.new.doit
dump_heap('heap3.jsonl')
