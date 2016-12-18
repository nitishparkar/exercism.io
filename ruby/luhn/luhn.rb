class Luhn
  attr_reader :digits

  def initialize(number)
    @digits = number.to_s.split('').map(&:to_i)
  end

  def addends
    output = []

    digits.reverse.each_with_index do |digit, i|
      output << digit && next if i.even?
      res = digit * 2
      if res < 10
        output << res
      else
        output << res - 9
      end
    end

    output.reverse
  end

  def checksum
    addends.inject(0, :+)
  end

  def valid?
    checksum % 10 == 0
  end

  def self.create(number)
    # Assume that the check digit is 0
    number = number * 10

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

end

