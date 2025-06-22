class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-abffad4895c3b0af709b7fe25544bb6885206998/example-go"
  version "abffad4895c3b0af709b7fe25544bb6885206998"

  def install
    bin.install "example-go"
    man1.install Dir["*.1"]
    bash_completion.install "example-go.bash" => "example-go"
    zsh_completion.install "example-go.zsh" => "_example-go"
    fish_completion.install "example-go.fish"
  end
end
