class Isogram
  def self.is_isogram?(string)
    chars_in_string = {}
    normalized = string.gsub(/[^[:alpha:]]/, '').downcase

    normalized.each_char do |c|
      return false if chars_in_string[c]
      chars_in_string[c] = true
    end

    true
  end
end

module BookKeeping
  VERSION = 2
end
