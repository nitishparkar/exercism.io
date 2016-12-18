class Luhn
  attr_reader :number

  def initialize(number)
    @number = number
  end

  def addends
    @addends ||= digits.reverse.each_with_index.map do |digit, position|
      position.even? ? digit : double(digit)
    end.reverse
  end

  def checksum
    addends.inject(0, :+)
  end

  def valid?
    check_digit.zero?
  end

  def self.create(partial_number)
    # Assume that the check digit is 0
    number = partial_number * 10

    # Calculate checksum with check digit 0
    # This works because we don't double the check digit in checksum calculations, just add it.
    checksum = Luhn.new(number).checksum


    # Get the right check digit. For example,
    # If checksum is 10 check digit should be 0
    # If checksum is 11 check digit should be 9
    # If checksum is 12 check digit should be 8, and so on...
    final_digit = (10 - checksum) % 10

    # We assumed that check digit was 0. Let's make that right
    number + final_digit
  end

  private

  def double(digit)
    res = digit * 2
    res < 10 ? res : res - 9
  end

  def check_digit
    checksum % 10
  end

  def digits
    @digits ||= number.to_s.chars.map(&:to_i)
  end
end

